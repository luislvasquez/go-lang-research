# GO RESEARCH

## Research Milestones:

### Milestone 1
* [x] Boot single endpoint
* [x] Containerize project
* [x] Add hot reload for dev
* [x] Configure backend API in seperate files
* [x] Extra: Look for pydantic alternative for q. params and payload

#### Summary
* `Go` local modules import was a nightmare to setup but makes sense after it
* `Gin` official documentation is uglier than latinoamerica's corruption stats
* Coming from `python`, syntax makes sense and shows huge potential

#### Resources:
* https://rpereira.pt/programming/setup-docker-for-go-development/
* https://developpaper.com/learn-the-validator-library-of-gin-frameworks-parameter-validation-this-article-is-enough/

### Milestone 2
* [x] Research async implementation approach
* [x] Use schema/model validation on queries
  * [x] Validate mandatory and optional body fields
* [ ] Try simple connection with a sqlite


#### Summary
* Async way to go: Go-rutines
* Basic validation is out of the box. Complex could be solved using [go-playground validator](https://github.com/go-playground/validator)