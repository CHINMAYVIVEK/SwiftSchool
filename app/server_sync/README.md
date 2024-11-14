# Server Sync
this package will sync the local data to server on demand

## Sync Process

- Get all table status from the server
- Match status from the local machine
- **Start the transaction**
- start the concurrent call to update table by table
- **If all Successful Commit**
- Check Table status of local and server
- **If transaction fails rollback**
- Show the Local vs Server report