Robert Cecil Martin:
> the database is a detail

> the framework is a detail

There are three layers.
1.  Buisness logic (entities and use cases)
2.  Storage. Now we just use a variable, but we can easy change to something else. For example, postgresql or third party api. **Because buisness logic don't depend from Storage.**
3. API. It is our external layer. Now it is Gin. Similarly we can use other web framework. Not even necessarily a web.

All in docker-compose :smile: that's why easy to run and to use.

*In this case we have union **Clear Architecture** and https://github.com/golang-standards/project-layout*
