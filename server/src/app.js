const express = require('express')
const bodyParser = require('body-parser')
const cors = require('cors')
const morgan = require('morgan')




const app = express()
app.use(morgan('combined'))
app.use(bodyParser.json())
app.use(cors())


var mysql = require('mysql');
//const mysql = require('promise-mysql');
//const mysql = require('mysql2/promise');

var pool = mysql.createPool({
    connectionLimit: 100,
    host: process.env.DATABASE_URL,
    user: 'root',
    password : process.env.MYSQL_ROOT_PASSWORD,
    database : process.env.MARIADB_DATABASE,
    debug: false
    
});






// SERVER Setup
app.get('/api/getdb/Personal/:range', (req, res) => {  
  
  // 'Range' wird bei Personal-Get nicht benötigt.
  // wurde aber der Einfachheit halber gelassen, da sonnst in der App ein 
  // extra DBGet gemacht werden müsste
  
   
  // sql = "select Abteilungen.id as id, Abteilungen.name as name , Firmen.name as firma, notifyMail, Abteilungen.enabled \
  // from Abteilungen \
  // left join Firmen on \
  // Abteilungen.firmaId = Firmen.id"

var sql = " select \
Personal.id as PersonalId, \
Firmen.name as firma, \
persnr, \
nachname,\
vorname, \
titel, \
geschlecht, \
eintritt, \
austritt, \
Abteilungen.name as abteilung, \
Person_Abteilung_Status.status as abteilungStatus, \
gebdat, \
dienstart, \
berufsbezeichnung, \
Person_Processor_Status.status as processorStatus \
from Personal \
left join Firmen on Firmen.id = Personal.firmaId \
left join Abteilungen on Abteilungen.id = Personal.abteilungId \
left join Person_Abteilung_Status on Personal.id = Person_Abteilung_Status.personId \
left join Person_Processor_Status on Personal.id = Person_Processor_Status.personId"

  console.log(sql)

  // console.log("sql-pwd: " + process.env.MYSQL_ROOT_PASSWORD )

  pool.getConnection(function(err, connection) {
    // do whatever you want with your connection here

    return  connection.query(sql, function (err, result, fields) {
      if (err) throw err
      else 
   //   console.log(result)
      res.json(result);
      connection.release();
    })       
  });
})

app.get('/api/getdb/Firmen/:range', async (req, res) => {  
   console.log("requested: /api/getdb/Firmen")
  // console.log("requested range: " + req.params.range)

  if (req.params.range == "enabled") {
    var sql ="Select * from Firmen where enabled = 1 ";
  } else {
    var sql ="Select * from Firmen order by enabled desc";
  }  
   
  console.log("sql: ",sql)

  // console.log("sql-pwd: " + process.env.MYSQL_ROOT_PASSWORD )

  pool.getConnection(function(err, connection) {    

    
  return  connection.query(sql, function (err, result, fields) {
      if (err){
        console.log("ERROR: ", err)
        throw err
      } 
      else 
      console.log(result)
      res.json(result);
      connection.release();
    })       
  });
})

app.get('/api/getdb/Abteilungen/:range', (req, res) => {  
  // console.log("requested: /api/getdb/Abteilungen")
  // console.log("requested range: " + req.params.range)

  // if (req.params.range == "enabled") {
  //   var sql ="Select * from " +  req.params.entrys + " where enabled = 1 ";
  // } else {
  //   var sql ="Select * from " +  req.params.entrys + " order by enabled desc";
  // }  
   
  sql = "select Abteilungen.id as id, Abteilungen.name as name , Firmen.name as firma, notifyMail, Abteilungen.enabled \
    from Abteilungen \
    left join Firmen on \
    Abteilungen.firmaId = Firmen.id"
  
  console.log(sql)

  // console.log("sql-pwd: " + process.env.MYSQL_ROOT_PASSWORD )

  pool.getConnection(function(err, connection) {
    // do whatever you want with your connection here

    return  connection.query(sql, function (err, result, fields) {
      
      if (err) throw err
      else 
    //  console.log(result)
      res.json(result);
      connection.release();
    })       
  });
})

app.get('/api/getdb/Processors/:range', (req, res) => {  
  // console.log("requested: /api/getdb/Processors")
  // console.log("requested range: " + req.params.range)  
   
  sql = "select Processors.id as id, Processors.name as name , Processors.name as firma, Processors.notifyMail, Processors.enabled \
    from Processors;"
  
  console.log(sql)

  // console.log("sql-pwd: " + process.env.MYSQL_ROOT_PASSWORD )

  pool.getConnection(function(err, connection) {
    // do whatever you want with your connection here

    return  connection.query(sql, function (err, result, fields) {
      
      if (err) throw err
      else 
    //  console.log(result)
      res.json(result);
      connection.release();
    })       
  });
})

app.get('/api/getdb/Items/:range', (req, res) => {  
  // onsole.log("requested: /api/getdb/Items")
  // console.log("requested range: " + req.params.range)
   
  sql = "select Items.id, Items.text, Items.type,Items.enabled,Abteilungen.name as depId, Processors.name as procId, Items.type \
    from Items  \
    left join Abteilungen on Abteilungen.id = Items.depId \
    left join Processors on Processors.id = Items.processorId;"
  
  console.log(sql)

  // console.log("sql-pwd: " + process.env.MYSQL_ROOT_PASSWORD )

  pool.getConnection(function(err, connection) {
    // do whatever you want with your connection here

    return  connection.query(sql, function (err, result, fields) {
    
      if (err) throw err
      else 
      console.log(result)
      res.json(result);
      connection.release();
    })       
  });
})
 


app.post('/api/addPersonal', async (req, res) => {
  var sql = "INSERT INTO `Personal` \
  (`firmaId`, \
  `persnr`,	\
  `nachname`,\
  `vorname`,\
  `titel`,\
  `geschlecht`,\
  `eintritt`,\
  `austritt`,\
  `abteilungId`,\
  `gebdat`,\
  `dienstart`,\
  `berufsbezeichnung`) \
  VALUES \
    ( ' " + req.body.perComp + " '," +
              "\
    ' " + req.body.perNr + " ',\
    ' " + req.body.perNachname + " ',\
    ' " +
              req.body.perVorname + " ',\
    ' " + req.body.perTitle + " ',\
    ' " + req.body.perSex +
              " ',\
    ' " + req.body.perEintritt + " ',\
    ' " + req.body.perAustritt + " '," +
              "\
    ' " + req.body.perAbt + " ',\
    ' " + req.body.perGebDat + " ',\
    ' " +
              req.body.perDienstart + " ',\
    ' " + req.body.perBeruf + " '); "

 pool.getConnection(async function(err, connection) {

  var delay = (seconds) => new Promise((resolves) =>{ setTimeout(resolves, seconds*1000)})
//console.log(connection)



 await connection.beginTransaction();

  var  newPersonId = null
  var _processors = null

    

            const insertNewPersonal = async () => Promise.resolve()
            
            connection.query(sql, function(err,result){
                if(err){
                  console.error("Error: " + err)
                  connection.rollback();
                  connection.release;
                }              
                console.log("inserted new Personal: " , result.insertId)              
                console.log("new Personal SQL: " ,sql)
                newPersonId = result.insertId;
                //console.log("newPersonId: " + newPersonId)               
              })
            
            
            
              sql2 = "insert into Person_Abteilung_Status (personId2,abteilungId,status) values (" + newPersonId + "," + req.body.perAbt +","+ 0 +");"
              console.log("newPersonId: " + newPersonId)  
              console.log("newPersonId sql: " + sql2)  
              connection.query(sql2, function(err,result){
               if(err){
                  console.log("ERROR Person_Abteilung_Status : \n", err.stack)
                  console.log("ERROR SQL:  ", sql2)
                  connection.rollback();
                }
                console.log("res  Person_Abteilung_Status:",result)
              });              
            
            
              sql2 = "insert into Person_Abteilung_Status (personId,abteilungId,status) values (" + newPersonId + "," + req.body.perAbt +","+ 0 +");"              
              console.log("newPersonId sql: " + sql2)  
              connection.query(sql2, function(err,result){
               if(err){
                console.log("ERROR Person_Abteilung_Status : \n", err.stack)
                console.log("ERROR SQL:  ", sql2)
                connection.rollback();
               }
                console.log("res  Person_Abteilung_Status:",result)
              });
            
            
              connection.query(' \
                  select  Processors.id, Processors.name from Personal \
                  left join Abteilungen on Abteilungen.id = Personal.abteilungId \
                  left join Items on Items.depId = Abteilungen.id \
                  right join Processors on Processors.id = Items.processorId \
                  where Processors.id in ( \
                  select  Processors.id from  Items \
                  right join Processors on Processors.id = Items.processorId \
                  where Processors.enabled=1 \
                  and Items.enabled=1     \
                  group by Processors.id \
                  ) and ( \
                  Abteilungen.id = Personal.abteilungId \
                   or abteilungId = 1 \
                  ) and \
                  Personal.id = ' + newPersonId + ';' , function (err,result) {
                  if (err) {
                    console.log("ERROR query Processors : \n", err.stack)                    
                    connection.rollback();
                  }
                 console.log("res select Processors:",result.name)
                _processors = result;
               });
            

             
              for(var i in _processors) {
                sql3 = 'insert into Person_Processor_Status(personId2,processorId,status) values ( '
                          + newPersonId + "," 
                          + result[i].id + "," + 0 + ');'
                          connection.query(sql3, function(err){
                            if (err) {
                              console.log("ERROR Person_Abteilung_Status : \n", err.stack)
                              console.log("ERROR SQL:  ", sql2)
                              connection.rollback();
                            }
                          })          
                }        
             
             
                connection.commit();
                connection.release();
              //res.json(result);
              res.send(400)
  
                     

           
  });
});




//////////////////////////////////

app.post('/api/addPersonal-test2', async (req, res) => {
    console.log("addPersonal got: ", req.body);

    // var sql = "INSERT INTO `appdb`.`Personal` (`name`, `enabled`) VALUES ('"
    // +req.body.data +"', '1');"

    var sql = "INSERT INTO `Personal` \
(`firmaId`, \
`persnr`,	\
`nachname`,\
`vorname`,\
`titel`,\
`geschlecht`,\
`eintritt`,\
`austritt`,\
`abteilungId`,\
`gebdat`,\
`dienstart`,\
`berufsbezeichnung`) \
VALUES \
  ( ' " + req.body.perComp + " '," +
            "\
  ' " + req.body.perNr + " ',\
  ' " + req.body.perNachname + " ',\
  ' " +
            req.body.perVorname + " ',\
  ' " + req.body.perTitle + " ',\
  ' " + req.body.perSex +
            " ',\
  ' " + req.body.perEintritt + " ',\
  ' " + req.body.perAustritt + " '," +
            "\
  ' " + req.body.perAbt + " ',\
  ' " + req.body.perGebDat + " ',\
  ' " +
            req.body.perDienstart + " ',\
  ' " + req.body.perBeruf + " '); "

    
    console.log("SQL Personal: " + sql)

    // https://stackoverflow.com/questions/47809269/node-js-mysql-pool-begintransaction-connection

    pool.getConnection(function (err, connection) {
         connection.beginTransaction(function (err) {
            if (err) { throw err; }             
              
                   // Person einfügen         
              connection.query(sql, function (err, result) {
                    if (err) {
                        connection.rollback(function () {
                            throw err;
                        });
                    }

                    var  newPersonId = result.insertId;
                    console.log("newPersonId: " + newPersonId)  

                    //// Abteilungs-Bearbeitungsstatus einfügen
                    sql = "insert into Person_Abteilung_Status (personId,abteilungId,status) values (" + newPersonId + "," + req.body.perAbt +","+ 0 +");"
                    console.log("SQL insert Person_Abteilung_Status : " + sql)
                    connection.query(sql, function (err, result) {
                      if (err) {
                          connection.rollback(function () {
                              throw err;
                          });                      
                      }
                      console.log(result)
                    });
                                            

                      // Bearbeiter / Processor bearbeitungsstatus einfügen.
                      //// SQL Select gibt alle Processoren zurück, die nach item-Abteilung-Prozessor Schema relevant sind 
                      //// 
                        // select  Processors.id, Processors.name from Personal
                        // left join Abteilungen on Abteilungen.id = Personal.abteilungId
                        // left join Items on Items.depId = Abteilungen.id
                        // right join Processors on Processors.id = Items.processorId
                        // where Processors.id in (
                        // 	select  Processors.id from  Items
                        // 	right join Processors on Processors.id = Items.processorId
                        //     where Processors.enabled=1
                        //     and Items.enabled=1    
                        // 	group by Processors.id
                        // ) and (
                        // 	Abteilungen.id = Personal.abteilungId
                        //     or abteilungId = 1
                        // )
                        ///////////////////////// 
                   
                    
                      connection.query(' \
                        select  Processors.id, Processors.name from Personal \
                        left join Abteilungen on Abteilungen.id = Personal.abteilungId \
                         left join Items on Items.depId = Abteilungen.id \
                         right join Processors on Processors.id = Items.processorId \
                         where Processors.id in ( \
                         	select  Processors.id from  Items \
                         	right join Processors on Processors.id = Items.processorId \
                             where Processors.enabled=1 \
                             and Items.enabled=1     \
                         	group by Processors.id \
                         ) and ( \
                         	Abteilungen.id = Personal.abteilungId \
                             or abteilungId = 1 \
                         ) and \
                         Personal.id = ' + newPersonId + ';' , function (err, result) {
                        if (err) {
                            connection.rollback(function () {
                                throw err;
                            });                      
                        }
                        for (var i in result) {
                          console.log('SQL-Results : ', result[i].name);

                          sql = 'insert into Person_Processor_Status(personId2,processorId,status) values ( '
                          + newPersonId + "," 
                          + result[i].id + "," + 0 + ');'
                          console.log ("inserting in Person_Abteilung_Status: \n" + sql)
                          connection.query(sql, function (err, result) {
                            if (err) {
                              
                                connection.rollback(function () {
                              //      throw err;                                        
                              console.log('SQL-Rollback: ', result);
                                });
                            }
                          })
                        }

                      
                      //  console.log(result)
                      });


                      /////////// TEST funktion
                      // connection.query('select * from Personal', function (err, result) {
                      //   if (err) {
                      //       connection.rollback(function () {
                      //           throw err;
                      //       });                      
                      //   }
                      //   for (var i in result) {
                      //     console.log('Post Titles: ', result[i].nachname);
                      // }
                      // //  console.log(result)
                      // });
                      res.send(200)

                })                                                 
                

                connection.commit(function(err) {
                  if (err) { 
                    connection.rollback(function() {
                      throw err;
                    });
                  }
                  console.log('Transaction Complete.');
                  connection.release();
                });
            
          });
        })
})

  


app.post('/api/addCompany', async(req,res)=> {  
    console.log("AddCompany got: ", req.body.data);  

    var sql = "INSERT INTO `appdb`.`Firmen` (`name`, `enabled`) VALUES ('" +req.body.data +"', '1');"
    pool.getConnection(function(err, connection) {

    return  connection.query(sql, function (err, result, fields) {
      if (err) throw err
      else 
      console.log(result)
      res.json(result);
      connection.release();
    })
  });  
  //res.status(201).send();
}) 
  

app.post('/api/AddItem', async(req,res)=> {  
  var sql = "INSERT INTO `appdb`.`Items` (`processorId`, `text`, `type`, `enabled`,`depId`) VALUES (" +req.body.itemProcessor +", '" +req.body.itemText +"','" +req.body.itemType +"', '1','" + req.body.itemDepartment+ "');"
  
  console.log("AddItem  requested SQL:  ",sql);  

  pool.getConnection(function(err, connection) {

  return  connection.query(sql, function (err, result, fields) {
    if (err) throw err
    else 
    console.log(result)
    res.json(result);
    connection.release();
  })
});   
//res.status(201).send();
})

app.post('/api/AddProcessor', async(req,res)=> {  
   console.log("AddNewProcessor got: ", req.body.data);  
   console.log("AddNewProcessor got: ", req.body.newProcessorName);  

    var sql = "INSERT INTO `appdb`.`Processors` (`name`,`notifyMail`, `enabled`) VALUES \
    ('" + req.body.newProcessorName +"', '" + req.body.newProcessorMail +"', '1');"

    console.log("AddNewProcessor requested SQL: ", sql);  
  
  

    pool.getConnection(function(err, connection) {

    return  connection.query(sql, function (err, result, fields) {
      if (err) throw err
      else 
      console.log(result)
      res.json(result);
      connection.release();
    })
  });  

}) 


app.post('/api/AddNewDep', async(req,res)=> {  
  // console.log("AddNewDep got: ", req.body.data);  
  // console.log("AddNewDep got: ", req.body.data.newDepName);  

    var sql = "INSERT INTO `appdb`.`Abteilungen` (`name`,`FirmaId`,`notifyMail`, `enabled`) VALUES \
    ('" + req.body.data.newDepName +"', '" + req.body.data.newDepCompany +"',  '" + req.body.data.newDepMail +"',    \
    '1');"

    console.log("AddNewDep requested SQL: ", sql);  
  
  

    pool.getConnection(function(err, connection) {

    return  connection.query(sql, function (err, result, fields) {
      if (err) throw err
      else 
      console.log(result)
      res.json(result);
      connection.release();
    })
  });  

}) 


app.post('/api/toggle/:entry/:id/:enabled', async(req,res) => {
  // console.log("toggle entry: " + req.params.entry + " : id : " + req.params.id)
  // console.log("company current enabled: " + req.params.enabled)
  if (req.params.enabled == 1) {
    toggle = 0;  
  } else  {
    toggle = 1;  
  }
  
  var sql = "update `appdb`."+ req.params.entry + " set `enabled` = '"+ toggle +"' where id = " +  req.params.id +";"

  console.log("toggle entry got SQL: " + sql)

  pool.getConnection(function(err, connection) {

    return  connection.query(sql, function (err, result, fields) {
      if (err) throw err
      else 
      console.log("SQL-Query Result: " + result.message)
      res.json(result);
      connection.release();
    })
  });  
})



// Default Route
app.get('/', function(req, res){
  console.log("handling default route - requested: " + req)
  res.status(404).end();
});


app.listen(process.env.PORT || 8081, () => {
  console.log("App Server started");
})
