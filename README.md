# Dining Actors Problem

Modelado y solución al problema de la cena de los filósofos usando actores, en Go. Se resuelve utilizando el algoritmo de Chandy/Misra

1. Para cada par de filosofos adyacentes, se crea un palito sucio asignado al filosofo con menor identificador (esto evita el deadlock)
2. Los palitos pueden estar sucios o limpios. Un filosofo solo puede comer si tiene ambos palitos limpios
3. Cuando un filosofo quiere un palito, debe solicitarselo al filosofo correspondiente.
4. Cuando un filosofo recibe un pedido de un palito, conserva el palito si esta limpio, y lo limpia y lo entrega si esta sucio.
5. Despues de que un filosofo haya terminado de comer, todos sus palitos quedan sucios. Si otro filosofo habia solicitado un palito previamente, entonces lo limpia y lo envía