**Diminuir uso de memória:**  
Ao armazenar uma _string_ de 10.000 caracteres, são necessários **10.000 bytes (≈ 9,77 KB ou 0,0095 MB)** de memória.  
Quando utilizamos um ponteiro para, por exemplo, passar essa _string_ para outra função, evitamos a cópia de todo o conteúdo. Em vez disso, passamos apenas a **referência**, que ocupa **8 bytes** em arquiteturas de 64 bits.

uma forma simples de ver isso é com o seguinte código:
**Diminuir uso de memória:**  
Ao armazenar uma _string_ de 10.000 caracteres, são necessários **10.000 bytes (≈ 9,77 KB ou 0,0095 MB)** de memória.  
Quando utilizamos um ponteiro para, por exemplo, passar essa _string_ para outra função, evitamos a cópia de todo o conteúdo. Em vez disso, passamos apenas a **referência**, que ocupa **8 bytes** em arquiteturas de 64 bits.


