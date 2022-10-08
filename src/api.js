export default {
    install(app, options) {
        app.config.globalProperties.$api = {
            deleteList(list) {
                //redirect to the delete router endpoint
                //delete the list
                this._deleteList(list.id).then(() => {
                    //remove the deleted list from the lists array
                    app.config.globalProperties.$store.state.lists.filter(i => i.id !== list.id);

                    //redirect to the first list of the array
                    app.config.globalProperties.$router.push({name: "app", params: {listId: app.config.globalProperties.$store.state.lists[0].id, mode: "view"}});
                }).catch(statusCode => {
                    //check if the status code is 401
                    if (statusCode === 401) {
                        //popup the login page

                    } else if (statusCode === 404) {
                        //redirect to the app page
                        app.config.globalProperties.$router.push({name: "app", params: {listId: app.config.globalProperties.$store.state.lists[0].id, mode: "view"}});
                    }
                });
            },
            getLists() {
                //get the lists from the api
                this._getLists().then((lists) => {
                    //complete the values of the params and check if they're valid
                    const listId = lists.find(i => i.id === app.config.globalProperties.$route.params.listId) ? app.config.globalProperties.$route.params.listId : lists[0].id;
                    const mode = ["view", "edit"].includes(app.config.globalProperties.$route.params.mode) ? app.config.globalProperties.$route.params.mode : "view";

                    //redirect to the app page
                    app.config.globalProperties.$router.push({ name: "app", params: { listId, mode } });

                    //populate the list
                    app.config.globalProperties.$store.commit("setLists", lists);
                }).catch(statusCode => {
                    //check if the status code is 401
                    if (statusCode === 401) {
                        //popup the login page

                    }
                });
            },
            _getLists() {
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

        }
    },
    _deleteList(listId) {
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
    },
    _createList(list) {
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
    },
}
