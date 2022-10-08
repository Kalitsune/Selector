function _deleteList(listId) {
    //delete a list from the api
    return new Promise((resolve, reject) => {
        fetch(`/api/list/${listId}`, {
            method: 'DELETE'
        }).then(response => {
            if (response.ok) {
                resolve();
            } else {
                reject(response.status);
            }
        })
    })
}
function _createList(list) {
    //create a list on the api
    return new Promise((resolve, reject) => {
        fetch('/api/lists', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(list)
        }).then(response => {
            if (response.ok) {
                response.json().then(data => {
                    resolve(data);
                })
            } else {
                reject(response.status);
            }
        })
    })
}

function handleAuth(props) {
    return new Promise((resolve, reject) => {
        //popup the login page
        props.$store.commit("needLogin", true);

        //wait for the authentication to complete and then resolve the promise
        const unsubscribe = props.$store.watch(
            (state) => state.needLogin,
            (newValue, oldValue) => {
                if (newValue === false) {
                    unsubscribe();
                    resolve();
            }
        });
    })
}

export default {
    install(app, options) {
        let props = app.config.globalProperties
        props.$api = {
            createList(list) {
                //create a disabled list with the given name
                props.$store.commit("setLists", [...props.$store.state.lists, {name: list.name, id: 0}]);

                //create a list on the api
                _createList(list).then(list => {
                    props.$store.commit("activateList", list);
                }).catch(statusCode => {
                    //if the user is not authenticated, popup the login page
                    if (statusCode === 401) {
                        handleAuth(props).then(() => {
                            this.createList(list);
                        })
                    }
                });
            },
            deleteList(list) {
                const listId = list.id;
                //disable the button
                list.id = 0;

                //delete the list
                _deleteList(listId).then(() => {
                    //remove the deleted list from the lists array
                    props.$store.commit("setLists", props.$store.state.lists.filter(i => i.id !== list.id));

                    //check if the selected list has been deleted
                    if (props.$route.params.listId === listId) {
                        //redirect to the first element if it has been
                        props.$router.push({name: "app", params: {listId: props.$store.state.lists[0].id, mode: "view"}});
                    }
                }).catch(statusCode => {
                    //enable the button back
                    list.id = listId;

                    //if the user is not authenticated, popup the login page
                    if (statusCode === 401) {
                        //popup the login page
                        handleAuth(props).then(() => {
                            //get the lists again
                            this.deleteList(list);
                        });
                    } else if (statusCode === 404) {
                        //check if the selected list has been deleted
                        if (props.$route.params.listId === listId) {
                            //redirect to the first element if it has been
                            props.$router.push({name: "app", params: {listId: props.$store.state.lists[0].id, mode: "view"}});
                        }
                    }
                });

            },
            getLists() {
                //get the lists from the api
                return new Promise((resolve, reject) => {
                    fetch('/api/lists').then(response => {
                        if (response.ok) {
                            response.json().then(data => {
                                resolve(data);
                            })
                        } else {
                            reject(response.status);
                        }
                    })
                })
            },
            refreshLists() {
                //get the lists from the api
                props.$api.getLists().then((lists) => {
                    //complete the values of the params and check if they're valid
                    const listId = lists.find(i => i.id === props.$route.params.listId) ? props.$route.params.listId : lists[0].id;
                    const mode = ["view", "edit"].includes(props.$route.params.mode) ? props.$route.params.mode : "view";

                    //redirect to the app page
                    props.$router.push({ name: "app", params: { listId, mode } });

                    //populate the list
                    props.$store.commit("setLists", lists);
                }).catch(statusCode => {

                    //if the user is not authenticated, popup the login page
                    if (statusCode === 401) {
                        //popup the login page
                        handleAuth(props).then(() => {
                            //refresh the lists
                            this.refreshLists();
                        });
                    }
                });
            },
        }
    },
}
