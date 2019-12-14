# [Dailyrotate](https://github.com/yoannduc/dailyrotate) Hook for [Logrus](https://github.com/sirupsen/logrus) <img src="http://i.imgur.com/hTeVwmJ.png" width="40" height="40" alt=":walrus:" class="emoji" title=":walrus:" />

Daily rotate log file hook for [logrus](https://github.com/sirupsen/logrus).

## Usage

```g
package main

import (
    "github.com/sirupsen/logrus"
    logrusdailyrotate "github.com/yoannduc/logrus-dailyrotate"
)

func main() {
    log := logrus.New()

    hook, err := logrusdailyrotate.New("", 5, logrus.InfoLevel, &logrus.TextFormatter{})
    if err != nil {
        // Handle error your way
    }

    log.Hooks.Add(hook)

    log.Info("Ready to go ! :)")
}
```
