# Sprint 2 Overview
## Work Completed in Sprint 2
- Integrated front end and back end to perform a simple (login or profile component)
### Front End
- Added a calendar to the calendar page that we struggled for a long time with in Sprint 1.
- Added a profile page for a user. (club to come later)
- Implemented a registration page that routes from the login page. (no button in the navbar in header file so had to route through login) Only takes in valid inputs from the user such as only allowing UF emails.
- Added a simple cypress test.
- Added unit tests.

### Back End
- Added user searched functionality
- Added user addition to DB functionality
- Added user login functionality
# Testing

## Front End Unit Tests
- added unit tests to each component with functions (login, register, calendar, and profile)
- at the moment, most work- we are having trouble with out integration as described in the video


## Front End Cypress Test
Wrote Cypress test that involved an up-down counter with buttons, implemented the following checks to accompany:
1. Make sure counter initializes to zero every time
2. Makes sure it is possible to change the counter value to a specific number
3. Makes sure that the '-' button decrease the counter, and '+' button increases the counter
4. Makes sure that clicking makes a change event happen with the counter

## Back End Unit Tests

# Back End Documentation
- ## GET
 ### able to search and find existing users
func (a *App) GetUserByID(id string)
func (a *App) IdHandler
### able to return detailed issues within searching
- ## POST
### adding user to DB
 func (a *App) AddUserHandler <- gets passed JSON information
### login credentials posting
func (a *App) loginHandler <- passed username and password creds with JSON
- ## TESTING
### func HealthCheck returns plain text if API is running.

