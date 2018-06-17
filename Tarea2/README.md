# Analisis-Algoritmos
ISC 405

### Tarea #2

1.-En esta tarea va a aplicar uno o más algoritmos para el problema 2SAT. Aquí hay 6 diferentes
instancias 2sat1.txt, 2sat2.txt, 2sat3.txt, 2sat4.txt, 2sat5.txt y 2sat6.txt..
El formato del archivo es el siguiente. En cada caso, el número de variables y el número de cláusulas es
el mismo, y este número se especifica en la primera línea del archivo. Cada línea subsiguiente
especifica una cláusula a través de sus dos literales, con un número que indica la variable y un signo "-"
denota lógica "no". Por ejemplo, la segunda línea del primer archivo de datos es "-16808 75250", que
indica la cláusula ¬x16808∨x75250.

Su tarea consiste en determinar cuál de los 6 casos son satisfiable, y que son unsatisfiable. En la
respuesta , introduzca una cadena de 6 bits, donde el bit i-ésimo debe ser 1 si la instancia i-ésimo es
satisfiable, y 0 en caso contrario. Por ejemplo, si usted piensa que los 3 primeros casos son satisfiable y
el último 3 no son, entonces usted debe responder con 111000 en la respuesta.

DISCUSIÓN: Esta asignación es deliberadamente abierta, y se puede poner en práctica lo que el
algoritmo 2SAT desea. Por ejemplo, 2SAT reduce al cálculo de los componentes fuertemente conexos
de un gráfico adecuado (con dos vértices por variable y dos bordes dirigidos por cláusula, usted debe
pensar en los detalles). Esto podría ser una opción especialmente atractiva para alguno que ha
codificado un algoritmo SCC . Como alternativa, puede utilizar el algoritmo de búsqueda local
aleatorizado de Papadimitriou. Un tercer enfoque es a través de retroceso. Mencionemos el capítulo 9
del libro Dasgupta-Papadimitriou-Vazirani, por ejemplo, para más detalles.


