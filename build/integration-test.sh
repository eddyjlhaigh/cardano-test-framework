# Set up two servers
./build/run-server.sh 50005 &
./build/run-server.sh 50006 &

# Call both servers
./build/run-client.sh 50005
./build/run-client.sh 50006

# Kill servers
fuser -k 50005/tcp
fuser -k 50006/tcp