import axios from 'axios';
//import { resolve } from "url";



//const url = `http://docker:5000/api/posts/`;
const apiURL = window.location.protocol + "//"+ window.location.hostname +":5000/api"

// --> /api/getdb/:entrys

class DBService {

    static toggleEntry(table,id,enabled){        

        const callUrl = apiURL + "/toggle"
        return axios.post(`${callUrl}`,{            
            table,
            id,
            enabled
        },{
        headers: {
            'Content-Type': 'application/json'
        }
        });          
    }

    
    static DBGetCompanys() {        
          const callUrl = apiURL + "/Companys"
          return axios .get(callUrl) 
              .then((response) =>  {
                  console.log(response.data)
                  return response.data.Firmen
              })
   
    }
    
    static DBAddCompany(data){
        const callUrl = apiURL + "/Companys"
        return axios.post(callUrl, {
                name: data
              }, {
                headers: {
                    'Content-Type': 'application/json',
                }
              });
    }

    static DBGetDepartments() {        
        const callUrl = apiURL + "/Departments"
        return axios .get(callUrl) 
            .then((response) =>  {
                console.log(response.data)
                return response.data.Department
            })
 
  }
  
  static DBAddDepartment(data){
      const callUrl = apiURL + "/Departments"
      return axios.post(callUrl, {
        NewDepartment: data
            }, {
              headers: {
                  'Content-Type': 'application/json',
              }
            });
  }

  static DBAddProcessor(data){
    const callUrl = apiURL + "/Processors"
    return axios.post(callUrl, {
      NewDepartment: data
          }, {
            headers: {
                'Content-Type': 'application/json',
            }
          });
}
  




////////////////  -> OLD nodejs Server

    static DBget(entrys,range) {
    //     return new Promise(async (resolve, reject) => {
    //    // return new Promise( (resolve, reject) => {
    //         try {
    //             ///////////////-------------->>>>>>>>>>>>>>>> apiCompanyUrl falsch 
    //             const res = await axios.get(apiCompanyUrl + entrys +"/" + range);
    //             //const res =  axios.get(apiurl + entrys +"/" + range);
    //             const data = res.data;
    //             resolve(
    //                 data.map(dataGrid => ({
    //                     ...dataGrid
    //                 }))
    //             );
    //         }catch(err) {
    //             reject(err);
    //         }
    //     })
    }

 static DBgetPersonal() {

 }

 
    static AddPersonal(perComp,perNr,perAbt,perSex,perTitle,perNachname,perVorname,perGebDat,perEintritt,perAustritt,perDienstart,perBeruf){
        const url = window.location.protocol + "//"+ window.location.hostname +":8081/api/addPersonal"        
        return axios.post(url, {
            perComp,
            perNr,
            perAbt,
            perSex,
            perTitle,
            perNachname,
            perVorname,
            perGebDat,
            perEintritt,
            perAustritt,
            perDienstart,
            perBeruf
              }, {
                headers: {
                    'Content-Type': 'application/json',
                }
              });
    }

    /// OLD -> NodeJS Server
    // static AddCompany(data){
    //     const url = window.location.protocol + "//"+ window.location.hostname +":8081/api/addCompany"        
    //     return axios.post(url, {
    //             data
    //           }, {
    //             headers: {
    //                 'Content-Type': 'application/json',
    //             }
    //           });
    // }
 
    
    // static AddNewDep(data){
    //     const url = window.location.protocol + "//"+ window.location.hostname +":8081/api/AddNewDep"        
    //     return axios.post(url, {
    //             data
    //           }, {
    //             headers: {
    //                 'Content-Type': 'application/json',
    //             }
    //           });
    // }

 
    
    static AddItem(itemText, itemDepartment, itemProcessor, itemType) {
    const url = window.location.protocol + "//" + window.location.hostname + ":8081" +
            "/api/AddItem"

    return axios.post(url, {
        itemText,
        itemDepartment,
        itemProcessor,
        itemType
    }, {
        headers: {
            'Content-Type': 'application/json'
        }
    });
    }

}

export default DBService;

