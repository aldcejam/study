---
resume:
  - k8s
---

```dataview
table join(file.tags, ", ") as "Tags"
from ""
where length(file.tags) > 0
and none(file.tags, (t) => contains(t, "excalidraw"))
```
