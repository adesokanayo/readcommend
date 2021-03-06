# Readcommend

Readcommend is a book recommendation web app for the true book aficionados and disavowed human-size bookworms. It allows to search for book recommendations with best ratings, based on different search criteria.

# Instructions

The front-end single-page app has already been developed using Node/TypeScript/React (see `/app`) and a PostgreSQL database with sample data is also provided (see `/migrate.*`). Your mission - if you accept it - is to implement a back-end microservice that will fetch the data from the database and serve it to the front-end app.

- In the `service` directory, write a back-end microservice in the language of your choice (with a preference for Go, if you know it) that listens on `http://localhost:5000`.
- Write multiple REST endpoints, all under the `/api/v1` route, as specified in the `/open-api.yaml` swagger file.
- The most important endpoint, `/books`, must return book search results in order of descending ratings (from 5.0 to 1.0 stars) and filtered according to zero, one or multiple user selected criteria: author(s), genre(s), min/max pages, start/end publication date (the "era"). A maximum number of results can also be specified.
- It's OK to use libraries for http handling/routing and SQL (ie: query builders), but try to refrain from relying heavily on end-to-end frameworks (ie: Django) and ORMs that handle everything and leave little room to showcase your coding skills! ;)
- Write some documentation (ie: at the end of this file) to explain how to deploy and run your service.
- Keep your code simple, clean and well-organized.
- If you use Git during development (and we recommend you do!), please ensure your repo is configured as private to prevent future candidates from finding it.
- When you are done, please zip your entire project (excluding the `.git` hidden folder if any) and send the archive to us for review.
- Don't hesitate to come back to us with any questions along the way. We prefer that you ask questions, rather than assuming and misinterpreting requirements.
- You have no time limit to complete this exercise, but the more time you take, the higher our expectations in terms of quality and completeness.
- You will be evaluated mainly based on how well you respect the above instructions. However, we understand that you may have a life (some people do), so if you don't have the time to respect all instructions, simply do your best and focus on what you deem most important.

# Development environment

## Docker Desktop

Make sure you have the latest version of Docker Desktop installed, with sufficient memory allocated to it, otherwise you might run into errors such as:

```
app_1         | Killed
app_1         | npm ERR! code ELIFECYCLE
app_1         | npm ERR! errno 137.
```

If that happens, first try running the command again, but if it doesn't help, try increasing the amount of memory allocated to Docker in Preferences > Resources.

## Starting front-end app and database

In this repo's root dir, run this command to start the front-end app (on port 8080) and PostgreSQL database (on port 5432):

```bash
$ docker-compose up --build
```

(later you can press Ctrl+C to stop this docker composition when you no longer need it)

Wait for everything to build and start properly.

## Creating and seeding database tables

In another terminal window, run this command to create and seed the PostgreSQL database:

```bash
$ ./migrate.sh
```

## Connecting to database

During development, you can connect to and experiment with the PostgreSQL database by running this command:

```bash
$ ./psql.sh
```

To exit the PostgreSQL session, type `\q` and press `ENTER`.

## Accessing front-end app

Point your browser to http://localhost:8080

Be patient, the first time it might take up to 1 or 2 minutes for parcel to build and serve the front-end app.

You should see the front-end app appear, with all components displaying error messages because the back-end service does not exist yet.

## Deploying and running back-end microservice
To start the backend service , ensure the app is ready and database is running. You should see a message that database system is ready to accept connection from the logs when you start the front and database.

 Then , you can start the backend in another terminal window. 
 
Navigate to the readcommendapi folder.
```bash 
$ cd .. service/readcommendapi
```
 
 **Build**

```bash 
$ docker build . -t readcommendapi
```

 **Run**

```bash 
$ docker run -p 5000:5000 readcommendapi
```

Send "http://localhost:5000/api/vi/books"  on your browser and that should return all the books in the database.

Your app is ready for use. You can now reload http://localhost:8080 and you can get your book recommendations.

Thank you.

 **Test**

```bash 
$ go test -v ./test/...
```
