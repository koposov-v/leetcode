## Задача №112542. Точки и отрезки

### Условие:
Дано `n` отрезков на числовой прямой и `m` точек на этой же прямой.  
Для каждой из точек определите, **скольким отрезкам** она принадлежит.

Точка `x` считается принадлежащей отрезку с концами `a` и `b`, если выполняется условие:

```
min(a, b) ≤ x ≤ max(a, b)
```

---

### Входные данные:
- Первая строка содержит два целых числа `n` и `m`  
  (`1 ≤ n, m ≤ 10⁵`) — количество отрезков и точек соответственно.
- Следующие `n` строк: по два целых числа `a` и `b` — концы отрезков.
- Последняя строка содержит `m` целых чисел — координаты точек.

Все числа по модулю не превосходят `10⁹`.

---

### Выходные данные:
Выведите `m` чисел — для каждой точки количество отрезков, которым она принадлежит.

---

### Пример 1:

**Вход:**
```
3 2
0 5
-3 2
7 10
1 6
```

**Выход:**
```
2 0
```

---

### Пример 2:

**Вход:**
```
1 3
-10 10
-100 100 0
```

**Выход:**
```
0 0 1
```

---

### Подсказка:
Чтобы решить задачу эффективно при `n, m ≤ 10⁵`, рекомендуется использовать сортировку и **двойной проход** по отсортированным событиям (sweep line / line sweep) или **дерево отрезков / бинарный индекс**.
