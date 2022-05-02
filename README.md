# Go p2p file sharing

a peer to peer file sharing CLI App built in Golang std lib (no 3rd party packages)


![demo](https://user-images.githubusercontent.com/58612131/166171343-4cc26cfa-98c5-48ff-9ee2-f474ea2254bc.gif)

features :

- safe concurrent senders 
- graceful shutdown & cleanup, using os signals
- fault-tolerant receiver
- reply on success/fail
- faster file r&w using bufferd i/o


Message shape : 

- filename & extension , 16 bytes
- filesize   , 10 bytes 
- file data , {filesize} bytes
