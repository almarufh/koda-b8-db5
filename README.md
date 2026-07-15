# Contact Management


## ERD

```mermaid

erDiagram

contact {
    bigint Id PK
    string First_Name
    string Last_Name
    string Email 
    string Phone
    timestamp CreatedAt
    timestamp UpdatedAt
}

```