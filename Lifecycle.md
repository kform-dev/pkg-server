# lifecycle

```mermaid
 %%{init: { 'theme':'dark', 'sequence': {'useMaxWidth':true} } }%%
---
 title: Nephio R2 Architecture
 config:
   theme: neutral
---

flowchart TD

subgraph Lifecycle
    direction TB

    Deleted((____Deleted___))
    Draft((_____Draft_____))
    Proposed((____Proposed___))
    Published((___Published___))
    DeletionProposed((DeletionProposed))
end

    
%% link 0
Deleted -- Create (Create Branch) --> Draft
%% link 1
DeletionProposed -- Delete --> Deleted
%% link 2
DeletionProposed -- Update (Lifecycle: Publish) --> Published
%% link 3
Draft -- Delete --> Deleted
%% link 4
Draft -- Update (Lifecycle: Deletion proposed) --> DeletionProposed
%% link 5
Draft -- Update (Lifecycle: Proposed) --> Proposed
%% link 6
Proposed -- Delete --> Deleted
%% link 7
Proposed -- Update (Lifecycle: Draft) --> Draft
%% link 8
Proposed -- Update (Lifecycle: Publish) (Create TAG)--> Published
%% link 9
Proposed -- Update (Lifecycle: Deletion proposed) --> DeletionProposed
%% link 10
Published -- Update (Lifecycle: Deletion proposed) --> DeletionProposed



classDef Node fill:#111111;

class Deleted Node
class Draft Node
class Proposed Node
class Published Node
class DeletionProposed Node

linkStyle default interpolate linear

```