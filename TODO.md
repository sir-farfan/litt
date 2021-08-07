Config
> Create yaml file  
> Read and pass around  
> Validate config

Service
> CRUD for Tag  
>> Clear tables between unit tests  
> CRUD for Member  
> CRUD for Member tags

Usecase
> Tag  
>> Restrict to 40 chars as defined in the DB  
> Member  
> Member tags

Controller
> Endpoints

Model
> ~~DB object~~

Database
> Docker compose  
>> ~~MySQL~~  
>> ~~Clear settings~~

> Connect from Go  
>> Perform insert  
>> perform select

> Dockerfile
>> Push to registry

CI/CD
> Quality  
>> ~~Lint~~

> Build on server  
> Lint on server  
> Create docker on server  
>> Push to registry (same as above)

