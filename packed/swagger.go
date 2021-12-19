package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GD41jMpmAEJiDBwMhSXJ6anpxbpQ2m9rOL8vNAQVgbGhzL7Elc9CwoICAjQ0vP38ffXDdoUFGB0yTjgUmNH2rOlT7UjNR2uNLhO+vzZS2iyt0PT1oSLUSsuNR2Z1JMUsvVXatuvpUdXuv7SXKnaOdNSY1IC68uVqhJNCZ2TjpRMiLwoYZRw83lwmM3rG/9y/N5fXtfqJeGVVGpwVlvv7FlfnVUsP/kW/Tz6K/X8rOj5X3Vrb+2P4OxuW3TlUdOfPY+uJK15x6xtaGRwQJMzbMKkqDkxUyIDJs2cz77pR+TzT+Lc1fPez/uTc09GJ+zQVlH19CXP1B6ZSq1Zt2STgmhKXqFnYFqcVZrFAQfdMEmfwBBBlvVh7N49z1LEkt3uMhx7frnvpb7YOole/zlKHh9PsipmMzIyMPz/H+DNzsG7bSYXJyMDwy9mBgZYmDIwfEcLU3ZEmIKD8bHMvkSQbmQ1Ad6MTCLMiDhBNhkUJzCwrRFE4o0hhFHYnQIBAgz/HUMYGbA4jJUNJM/EwMTQycDAMIkRxAMEAAD//yIRdBExAgAA"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
