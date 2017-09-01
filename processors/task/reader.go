package task

unc (collector *Collector) Execute() <- chan Message {
var wg sync.WaitGroup
wg.Add(1)

out := make(chan Message)

go func () {
collector.collect(collector.name, collector.cfg, out)
wg.Done()
}()

go func () {
wg.Wait()
close(out)
}()

return out
}
