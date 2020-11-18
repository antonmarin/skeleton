# Title

## Description

Business process application presenting. Ex: Catalog, Shopper registration

Production URL: <https://domain.com>

Business documentation: <https://confluence.domain.com>

Specification: <https://swagger.domain.com/serviceA.yml>, <https://confluence.domain.com/EventBus>

## Strategy of project

MVP till 02.2020 event / Stable

### Basic risks

- Low knowledge of domain
- Must be in production before NY

### Principles

- Use design by domain / Fill map of business to application links / Analyse before develop
- Controller per business scenario + smoke test
- TDD + nodesign
- hexagon + split unit,integration tests
- acceptance tests at <https://github.com/>...

## Contacts

- Technical: Serg @serg
- Product: Ivan @ivan
- `#tech`. Review, design assistance.
- Report bugs at jira.domain.com/BUGS.
- Current operation status at <http://status.domain.com>.

## Staging

How to deploy for acceptance testing and demo here

## Troubleshooting

Catching Symfony Exception "Type error: " at ***,
deploy new dump.

To solve TransportException redeploy app <https://jenkins.domain.com/job/>
