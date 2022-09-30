
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
     }
}
