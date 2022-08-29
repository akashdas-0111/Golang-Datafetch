Golang Data fetch memory optimization and testing:

1.	Postgres setup and SQL
2.	Fetch Data from Postgres
3.	Client to receive data
4.	Memory optimization
5.	Analysis of memory 

1.	Postgres setup and SQL: Setting up postgres using Docker and connecting DBeaver (SQL Editor) to insert table and data.
2.	Fetch Data from Postgres: Fetching data from postgres database using golang “database/sql” and publishing them using “net/http” package.
3.	Client to receive data: Reading the content of a webpage using “net/http” package.
4.	Memory optimization: To reduce the constant memory retention of golang “runtime/debug” free up the space used, back to the OS this helps us to clear the RAM.
5.	Analysis of memory: PPROF is a golang tool which gives us a profile with which we can see the detailed information.

To see detailed information in <br>
Web: - <br>
•	localhost: ____/debug/pprof <br>
Download options: - <br>
•	localhost: ____/debug/pprof/profile <
•	localhost: ____/debug/pprof/heap
How to use pprof:
Steps:
1.	Use the cmd
2.	Go tool pprof ___(profile name)
3.	Get report

Analysis of Memory:
Without debug.FreeOSMemory:
![image](https://user-images.githubusercontent.com/61954560/186821976-66864488-4a1d-49c5-9cff-c41166d2cc2c.png)

 
Observation: The memory constantly is increasing and at time of requesting the response the memory increased drastically and it 
Keep on increasing. The Garbage collection is also zero.




With debug.FreeOSMemory: 

![image](https://user-images.githubusercontent.com/61954560/186822014-0e4a1105-2e9a-4667-9ccd-198911b31b78.png)

 
Observation: The memory constantly is increased and after that it decreased and at the time of response it increased due to the store and fetch of data after that the fetch didn’t increase much .

When the DB has 1 billion and fetch data from the DB.
-	The golang database/sql is used to access the database and then open the DB using Query we put the pointer and fetch the rows one by one using rows.Next .

Before Optimization 

![image](https://user-images.githubusercontent.com/61954560/186822074-2836e0a6-dbe2-4d4e-ae51-d50bab02d1a8.png)


After Optimization

![image](https://user-images.githubusercontent.com/61954560/186822108-e886a580-ce53-40a1-b742-6a805b6d9943.png)

 
•	Why Postgres?
Postgres is used in industries because it gives the ability to use the database regardless of quantity of data. It also offers a lot of functionality to work.<br>
•	Why Golang?
It is better because managing thread in golang is easy using go routines and this property thing made with golang highly scalable and the garbage collection is also useful for memory.
•	Ran out of memory error
