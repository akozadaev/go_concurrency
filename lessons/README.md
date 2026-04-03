# Курс: конкурентность в Go

Рекомендуемый **порядок изучения** - по номерам папок. Сначала модель выполнения и каналы, затем контекст, затем примитивы **`sync`** (от гонки к `Pool`), в конце - **`errgroup`**.

| № | Папка | Тема |
|---|--------|------|
| 01 | [`01-goroutines`](01-goroutines) | Горутины; завершение `main` без ожидания дочерних задач. |
| 02 | [`02-channels-wait`](02-channels-wait) | Канал "готово", буфер, ожидание двух горутин. |
| 03 | [`03-channels-range`](03-channels-range) | `close`, `for range` по каналу. |
| 04 | [`04-context`](04-context) | Таймаут, отмена, `WithValue` (три файла - см. README). |
| 05 | [`05-sync-race`](05-sync-race) | Data race на `map`. |
| 06 | [`06-sync-mutex`](06-sync-mutex) | `sync.Mutex` + `map`. |
| 07 | [`07-sync-rwmutex`](07-sync-rwmutex) | `sync.RWMutex`, читатели и писатель. |
| 08 | [`08-sync-atomic-int64`](08-sync-atomic-int64) | `atomic.Int64`. |
| 09 | [`09-sync-atomic-funcs`](09-sync-atomic-funcs) | `atomic.AddInt64` / `LoadInt64`. |
| 10 | [`10-sync-map`](10-sync-map) | `sync.Map`. |
| 11 | [`11-sync-once`](11-sync-once) | `sync.Once`, `WaitGroup`. |
| 12 | [`12-sync-cond`](12-sync-cond) | `sync.Cond`. |
| 13 | [`13-sync-pool`](13-sync-pool) | `sync.Pool`. |
| 14 | [`14-errgroup`](14-errgroup) | `errgroup.WithContext`. |

**Задания для самостоятельной работы** - в [`EXERCISES.md`](EXERCISES.md).

**Продвинутые задачи по `sync` (методичка, 10 заданий с верификацией)** - [`SYNC_ADVANCED_TASKS.md`](SYNC_ADVANCED_TASKS.md). Эталонные решения - локально в `../solutions_advanced_sync/` (см. конец методички; каталог в `.gitignore`).

Общий обзорный конспект (модель планировщика, `select`, таблицы) - в [корневом README.md](../README.md).
