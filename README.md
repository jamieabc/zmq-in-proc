This repository is for an example of using zeromq as inter-process communication.

zeromq [PAIR](http://zguide.zeromq.org/page:all#Messaging-Patterns) mode is exclusive, which means it limits sender and receiver to be 1-1 relationship, which can gurantee a message will not be divided due to load balance. 
