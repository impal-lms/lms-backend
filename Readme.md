# Backend LMS

## Packages Docs
### Domain
should contains all the data models and it's business logic and validation

### Repository
repository is in charge of storage no  other logic should be in here other than update, create, and delete into storage. each subfolder is the implementation based on  database used, mock package is used only for early development for testing purpose

### Services
services used as a wrapper between domain and repository, containing logic that will be used in application

### Handler
as the name suggest, the http handler, responsible processing the http  request and response

## Migration  step
1. Create a new migration file on `migration/`
2. the migration need to be  following the example  on create_user
3. copy `.env.example` as `.env`
3. run `make migrate`