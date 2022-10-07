
export default {
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
     deleteList(listId) {
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
}
