Feasibility to execute tink-cli commands from the host without using docker-exec. 

Problem: 
1) User can only execute tink-cli commands only form the container.
2) User can only execute tink-cli commands using docker exec.

solution:
1) User should be authenticated before executing the command. 
2) User can execute tink-cli commands from the host or other node without using docker-exec. 
3) User, once authenticated, should be able to execute tink-cli commands from outside the provisioner.
4) Communication between the host (from where the cmd is executed) and the provisioner container should be secured.  
