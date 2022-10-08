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
        fetch('/api/list', {
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
            deleteList(list) {
                const listId = list.id;
                //disable the button
                list.id = 0;

                //delete the list
                _deleteList(listId).then(() => {
                    //remove the deleted list from the lists array
                    props.$store.commit("setLists", props.$store.state.lists.filter(i => i.id !== list.id));

                    //redirect to the first list of the array
                    props.$router.push({name: "app", params: {listId: props.$store.state.lists[0].id, mode: "view"}});
                }).catch(statusCode => {
                    //enable the button back
                    list.id = listId;

                    //check if the status code is 401
                    if (statusCode === 401) {
                        //popup the login page
                        handleAuth(props).then(() => {
                            //get the lists again
                            this.deleteList(list);
                        });                    } else if (statusCode === 404) {
                        //redirect to the app page
                        props.$router.push({name: "app", params: {listId: props.$store.state.lists[0].id, mode: "view"}});
                    }
                });

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
                    //check if the status code is 401
                    if (statusCode === 401) {
                        //popup the login page
                        handleAuth(props).then(() => {
                            //get the lists again
                            this.getLists();
                        });
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
            }
        }
    },
}
