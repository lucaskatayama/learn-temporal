# Notification

A notification system made with Temporal

```mermaid
stateDiagram-v2
    [*] --> GetInfo
    
    GetInfo --> Web
    GetInfo --> iOS 
    GetInfo --> Android
    GetInfo --> Email
    
    state join_state <<join>>
    
    Web --> join_state : result
    iOS --> join_state : result
    Android --> join_state : result
    Email --> join_state : result
    
    join_state --> Record
    
    Record --> [*]
    
```
