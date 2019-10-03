<template>

    <div class="ManageItems">
        
       <!--      <h1> {{msg}} </h1> -->
              <section class="form">
            <input class="form-control" v-model="newItemText" placeholder="Item-Text"/>
            Abteilung, für die das Item verfügbar ist ( ALL = Alle Abteilungen): 
            <select name="NewItemDepartment" v-model="selectedDepartment" @click=updateAllData() class="btn btn-secondary dropdown-toggle" >
               <!--  <option :value="comp.id">{{comp.id}} </option> -->
               <option v-for="dep in Abteilungen" v-bind:key="dep.id" :value="dep.id">{{dep.name}} </option>               
           </select> <br>
           Bearbeitende Abteilung:
           <select name="NewItemProcessor" v-model="selectedProcessor" @click=updateAllData() class="btn btn-secondary dropdown-toggle" >
                <!--  <option :value="comp.id">{{comp.id}} </option> -->                
                <option v-for="proc in Processors" v-bind:key="proc.id" :value="proc.id">{{proc.name}} </option>               
           </select><br>
           Typ des Items:
           <select name="NewItemType" v-model="selectedType" class="btn btn-secondary dropdown-toggle" >
                <!--  <option :value="comp.id">{{comp.id}} </option> -->                
                <option >CheckBox</option>               
                <option >Choose</option>               
                <option >Text</option>               
           </select> <input v-if="selectedType === 'Choose'" class="form-control" v-model="newItemChooseOptions" placeholder="Optionen, die zur Auswahl stehen. ( Eins;Zwei;Drei; )"/>

           <br>
            <button type="button" class="btn btn-success" @click='AddItem'>Item anlegen</button>

            
            <div class="entryRow form-group">
                <table class="table  table-striped table-dark" >
                   <thead class="thead-dark">
                      <tr>
                        <th scope="col">Item</th>
                        <th scope="col">Abteilung</th>
                        <th scope="col">Bearbeiter</th>
                        <th scope="col">Type/Value</th>
                        <th scope="col">(de-)aktivieren</th>
                      </tr>
                    </thead>
                    <tbody >
                    <tr v-for="Item in Items" v-bind:key="Item.id"> 
                        <template v-if="Item.enabled === 1">
                            <td>{{Item.text}}</td>                        
                            <td>{{Item.depId}}</td> 
                            <td>{{Item.procId}}</td> 
                            <td>{{Item.type}}</td> 
                            <td align = "right">    <button type="button" class="btn-disable" @click='toggleItem(Item.id, Item.enabled)'>disable</button> </td>
                        </template>
                        <template v-if="Item.enabled === 0">
                            <td class="disabledColumn">{{Item.text}}</td>
                            <td class="disabledColumn">{{Item.depId}}</td> 
                            <td class="disabledColumn">{{Item.procId}}</td> 
                            <td class="disabledColumn">{{Item.type}}</td> 
                            <td class="disabledColumn" align = "right">    <button type="button" class="btn-enable" @click='toggleItem(Item.id, Item.enabled)'>Enable</button> </td>
                        </template>                        
                    </tr>
                    </tbody>
                </table>
            </div>
              </section>
              

    </div>
</template>

<script>
import DBService from '../DBService';
//import TodoItem from './TodoItem.vue';
export default {
    name: 'ManageItems',
    components :{    
    },
    data() {
        return {
            Items:[],                    
            newItemText: null,
            Abteilungen:[],
            Processors:[],
            selectedDepartment: "",
            selectedProcessor: "",
            selectedType: "",
            newItemChooseOptions: "",

        }        
    },
    props: {
        msg: String 
    },
     created() {
        this.updateAllData();
  
    },
    methods: {
        async updateAllData() {
            this.Items = await DBService.DBget("Items","all"); 
            this.Abteilungen = await DBService.DBget("Abteilungen","all");            
            this.Processors = await DBService.DBget("Processors","all");     

       
        },
        async AddItem() {       
            if  (!this.selectedDepartment)  this.selectedDepartment = 0;

            if (this.selectedType == "Choose") {
                this.selectedType = "Choose:" + this.newItemChooseOptions;
            }

            if (this.selectedType == "" || this.selectedProcessor == "") {
                console.log("Error: no type or Processor choosen")
            } else {
                await DBService.AddItem(this.newItemText,this.selectedDepartment, this.selectedProcessor, this.selectedType);
            }
            
            
            this.newItemText = "";            
            this.selectedDepartment = "";
            this.selectedProcessor = "";
            this.selectedType = "";
            this.newItemChooseOptions = "";

            this.updateAllData()
        },

        async toggleItem(id,enabled){            
            await DBService.toggleEntry("Items",id,enabled);
        this.updateAllData()
        }
    }
            
    
}
</script> 


                                <!-- Add "scoped" attribute to limit CSS to this component only -->
                                <style scoped="scoped">
                                    div.container {
                                        max-width: 90%;
                                        
                                    }

                                  
                                    p.error {
                                        border: 1px solid #ff5b5f;
                                        background-color: #ffc5c1;
                                        padding: 10px;
                                        margin-bottom: 15px;
                                    }

                                    div.entryRow {
                                        position: relative;
                                        border: 1px solid #5bd658;
                                        background-color: #bcffb8;
                                        padding: 10px 10px 1px;
                                        margin-bottom: 0px;
                                    }

                                    div.created-at {
                                        position: absolute;
                                        top: 0;
                                        left: 0;
                                        padding: 5px 15px;
                                        background-color: darkgreen;
                                        color: white;
                                        font-size: 13px;
                                    }

                                    p.text {
                                        font-size: 22px;
                                        font-weight: 700;
                                        margin-bottom: 0;
                                    }
                                    
                                    p.smalltext {
                                        font-size: 14px;
                                        font-weight: 500;
                                        margin-bottom: 0;
                                    }
                                    /* table {
                                      width: 100%;
                                      text-align: left;
                                      border: 1px solid black;
                                      border-collapse: collapse;
                                           
                                    }
                                    th {
                                      padding: 5px;
                                      font-weight: 700;
                                      border: 1px solid black;
                                      vertical-align: middle;
                                    } */
                                     td {
                                       padding: 4px;
                                       font-size: 14px;
                                       font-weight: 500;
                                       margin-bottom: 0;
                                       border: 1px solid black;
                                       vertical-align: middle;
                                       align: left;
                                    }
                                    .btn-disable {
                                        background-color: rgb(53, 53, 53);
                                        color: rgb(236, 48, 48);
                                        border: 2px solid rgb(179, 7, 7); /* Green */
                                    }
                                    .btn-enable {
                                        background-color: rgb(53, 53, 53);
                                        color: rgb(124, 223, 148);
                                        border: 2px solid rgb(0, 145, 43); /* Green */
                                    }
 
                                </style>