# Trie Data Structure

A trie is a tree structure which is to store words. Each node represent a part of a word or a word.
The path from the root to leaf node we build some words.

## Trie tree Example:

[![](https://mermaid.ink/img/pako:eNp9Vd1u0zAYfRWTaVILacua0a0ZK4KNCyQYaNsNanfhxF8aa64dbIeuqnrNU_AOE7dIPEO543Fw7KRJikQuIsfnnO_X_rL2YkHAC725xFmKbt_MODKPFEJfGaDTQajbRb3eBEXuO-p22xQLKgeqCnTvaIdDqYaKUEPUIfRfJHdI3jYKNQE7Aq4IDYhVHllbjRsMnFwXX1N512SwlhWcvDcvuzNlLR5tsS4tg7QYeYvx2TJWLYaqC6T2ClRDuoT0Xh1Uw7xqJFulofbTULssHEXXDFE5EZUTodv6T7wQZzvx4SFSX6Y3X3IsAakUZ3BnuTHtdC6ojFm5W0dUKPLInbPXuy3ri0w-ElJa6fXQ7VIgRjm8jORgAmQOKBaLBXBtPUjREhO6vqR4IThBS6pTZEVWjiIJ-H6Den2r6lyLnBMghjBRNm67skHuWTw_n0gxrARINbKsqcBJI7UroWkMSKdYIy6QhgeNKHcaVKgxKUylYJaUKw2YODJVCGcZWEdJLrVhICKWteWyQySYXANmbIWY4HPnwCZcpGozLdIxLjFHdTUbEV64Ghp3ibY-8sg0KQMZF5VVdM5V7bPTece5YQ1MQ4tmFraxKbESCzCxEGRrl0FMMUNxiiWOjVFVzoqk4_dfPXnWe_rnR6P98UpOL1aSMkZj0-aJ-R62Dwvaft8-_v62fdz-3P7qds8qLYoZVuoSEjSXABwlxkZ4ME5GvtJS3EN4EARBue4tKdFpOMwezvbFQmJuDpNTJ-P_qI_31OYI-NXA8-008O318N2w8O0w8O2F9-2I8csx4rtb5-uS7hJomSa0DGzGPd9bgFxgSsxEXhesmWcOxAJmXmiWBBKcMz3zZnxjqHlGsIa3hGohvTDBTIHv4VyLmxWPvVDLHCqSuSHm4i12LCYwASNae3qV2fFPlTYmY8ETOi_2c8nMdqp1psLBoID7c3Pa8qhv7uJAUWJ6rtOv49FgNByd4mEAo5MAvwgCEkdH49NkeHyUkJPnR0PsbTa-BzbKD-5fY385m79dNQw9?type=png)](https://mermaid.live/edit#pako:eNp9Vd1u0zAYfRWTaVILacua0a0ZK4KNCyQYaNsNanfhxF8aa64dbIeuqnrNU_AOE7dIPEO543Fw7KRJikQuIsfnnO_X_rL2YkHAC725xFmKbt_MODKPFEJfGaDTQajbRb3eBEXuO-p22xQLKgeqCnTvaIdDqYaKUEPUIfRfJHdI3jYKNQE7Aq4IDYhVHllbjRsMnFwXX1N512SwlhWcvDcvuzNlLR5tsS4tg7QYeYvx2TJWLYaqC6T2ClRDuoT0Xh1Uw7xqJFulofbTULssHEXXDFE5EZUTodv6T7wQZzvx4SFSX6Y3X3IsAakUZ3BnuTHtdC6ojFm5W0dUKPLInbPXuy3ri0w-ElJa6fXQ7VIgRjm8jORgAmQOKBaLBXBtPUjREhO6vqR4IThBS6pTZEVWjiIJ-H6Den2r6lyLnBMghjBRNm67skHuWTw_n0gxrARINbKsqcBJI7UroWkMSKdYIy6QhgeNKHcaVKgxKUylYJaUKw2YODJVCGcZWEdJLrVhICKWteWyQySYXANmbIWY4HPnwCZcpGozLdIxLjFHdTUbEV64Ghp3ibY-8sg0KQMZF5VVdM5V7bPTece5YQ1MQ4tmFraxKbESCzCxEGRrl0FMMUNxiiWOjVFVzoqk4_dfPXnWe_rnR6P98UpOL1aSMkZj0-aJ-R62Dwvaft8-_v62fdz-3P7qds8qLYoZVuoSEjSXABwlxkZ4ME5GvtJS3EN4EARBue4tKdFpOMwezvbFQmJuDpNTJ-P_qI_31OYI-NXA8-008O318N2w8O0w8O2F9-2I8csx4rtb5-uS7hJomSa0DGzGPd9bgFxgSsxEXhesmWcOxAJmXmiWBBKcMz3zZnxjqHlGsIa3hGohvTDBTIHv4VyLmxWPvVDLHCqSuSHm4i12LCYwASNae3qV2fFPlTYmY8ETOi_2c8nMdqp1psLBoID7c3Pa8qhv7uJAUWJ6rtOv49FgNByd4mEAo5MAvwgCEkdH49NkeHyUkJPnR0PsbTa-BzbKD-5fY385m79dNQw9)

## Time Complexity

The time complexity of a trie for the operations, insert and delete over a string w with length n is O(n). In a Trie we have to follow at most n pointers from the root to the node wich is has isEnd true, or, if the word w exists in the trie. For each pointer, we expend O(1) for lookup up the correct pointer. So the overall time complexity is: 

```
TimeComplexity = n * O(1) = O(n) 
```

The time complexity is independent of the number of words stored in the trie and depends only of the length of the word k.
