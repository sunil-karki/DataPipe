# AirFlow PipeLine
----------------------------------------------------------

    Contains Nodes/DAGs for ETL Process

**PipeLine SetUp and StartUp**

**Step 1:**     
Follow the link https://towardsdatascience.com/getting-started-with-airflow-using-docker-cd8b44dbff98
before initial setup and also install docker in the local machine.

See Docker Images by
    
    docker images
or
    
    docker ps

----------------------------------------------------------

**Step 2:** 
Run Airflow UI by command:

    docker run -d -p 8080:8080 puckel/docker-airflow webserver

----------------------------------------------------------

**Step 3:** 
Check this link: 

    http://localhost:8080/admin/

----------------------------------------------------------

**Step 4:** 
Jump into running container’s command line using the command:

    docker exec -ti <container name> bash

_Here inside the container, the path /usr/local/airflow/ contains logs and dags
for the Airflow._

----------------------------------------------------------

**Step 5:** 
Run the Command:

    docker run -d -p 8080:8080 -v /home/localPath/PipeLine/:/usr/local/airflow/dags puckel/docker-airflow
The DAG created in the LocalPath _/home/localPath/PipeLine/_ will be mapped to the path _/usr/local/airflow/dags_ 
inside the docker container named : _puckel/docker-airflow_

A DAG 'Helloworld' is created in the main.py

-----------------------------------------------------------
-----------------------------------------------------------

**PipeLine Stop**

    docker stop <container_name>

-----------------------------------------------------------

**PipeLine Remove**
    
    docker rmi <container_name_1> <container_name_2>

-----------------------------------------------------------
-----------------------------------------------------------

_Helpful Links:_

   https://medium.com/@jacksonbull1987/how-to-install-apache-airflow-6b8a2ae60050

   https://towardsdatascience.com/getting-started-with-airflow-using-docker-cd8b44dbff98

   https://github.com/vishalsatam/Data-Pipelining/tree/master/Airflow
