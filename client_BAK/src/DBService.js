import axios from 'axios';
//import { resolve } from "url";



//const url = `http://docker:5000/api/posts/`;
const apiurl = window.location.protocol + "//"+ window.location.hostname +":8081/api/getdb/"
// --> /api/getdb/:entrys

class DBService {


    static DBget(entrys,range) {
        return new Promise(async (resolve, reject) => {
            try {
                const res = await axios.get(apiurl + entrys +"/" + range);
                const data = res.data;
                resolve(
                    data.map(dataGrid => ({
                        ...dataGrid
                    }))
                );
            }catch(err) {
                reject(err);
            }
        })
    }



    static AddCompany(data){
        const url = window.location.protocol + "//"+ window.location.hostname +":8081/api/addCompany"        
        return new Promise(async (resolve, reject) => {
            const res = await axios.post(url, {
                data
              }, {
                headers: {
                    'Content-Type': 'application/json',
                }
              })
              .then(function (response) {                                                                                                
               resolve(response);
              })
              .catch(function (error) {
                
              });
            })
    }
 
    
    static AddNewDep(data){
        const url = window.location.protocol + "//"+ window.location.hostname +":8081/api/AddNewDep"        
        return new Promise(async (resolve, reject) => {
            const res = await axios.post(url, {
                data
              }, {
                headers: {
                    'Content-Type': 'application/json',
                }
              })
              .then(function (response) {                                                                                                
               resolve(response);
              })
              .catch(function (error) {
                
              });
            })
    }

    static AddProcessor(newProcessorName, newProcessorMail){
        const url = window.location.protocol + "//"+ window.location.hostname +":8081/api/AddProcessor"        
        return new Promise(async (resolve, reject) => {
            const res = await axios.post(url, {
                newProcessorName,
                newProcessorMail
              }, {
                headers: {
                    'Content-Type': 'application/json',
                }
              })
              .then(function (response) {                                                                                                
               resolve(response);
              })
              .catch(function (error) {
                
              });
            })
    }

    static AddItem(itemText, itemDepartment, itemProcessor, itemType){
        const url = window.location.protocol + "//"+ window.location.hostname +":8081/api/AddItem"        
        return new Promise(async (resolve, reject) => {
            const res = await axios.post(url, {
                itemText,
                itemDepartment,
                itemProcessor,
                itemType
              }, {
                headers: {
                    'Content-Type': 'application/json',
                }
              })
              .then(function (response) {                                                                                                
               resolve(response);
              })
              .catch(function (error) {
                
              });
            })
    }


    static toggleEntry(table,id,enabled){
        const url = window.location.protocol + "//"+ window.location.hostname +":8081/api/toggle/"+ table +"/"
        
        //return axios.post(url + "/:"+ id + "/:"+ enabled)            
        return axios.post(`${url}${id}/${enabled}`)            
    }

    // static toggleCompay(id,enabled){
    //     const url = window.location.protocol + "//"+ window.location.hostname +":8081/api/toggleCompay/"
        
    //     //return axios.post(url + "/:"+ id + "/:"+ enabled)            
    //     return axios.post(`${url}${id}/${enabled}`)            
    // }
    // static toggleDepartment(id,enabled){
    //     const url = window.location.protocol + "//"+ window.location.hostname +":8081/api/toggleDepartment/"
        
    //     //return axios.post(url + "/:"+ id + "/:"+ enabled)            
    //     return axios.post(`${url}${id}/${enabled}`)            
    // }
}

export default DBService;

