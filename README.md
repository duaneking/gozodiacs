# GoZodiacs

Golang/Go lib to calculate zodiac signs for both Western (monthly) and Chinese (yearly) astrology.

A simple module in GoLang that handles multiple zodiac types correctly without the bad/dead 
code or poor test posture of other zodiac libraries available.

# Why
I needed a simple topic I could publish a "go module" on that I could read the wikipedia page of, understand, and fully Implement, and give 100% code coverage to.. Without it being something that a company would consider a huge enterprise problem that would get political or be fought over. I honestly just want to show how to do code coverage.

# Unit Test Coverage
This module has 100% test coverage using standard "go" tools and uses no non-system level depdencies, by design.

I often see modules published that do not meet this low and simple quality bar, but I'm not willing to publish a module myself here that violates that quality bar, especually for such a simple task.

# Supported Types of Zodiacs
For "Buisness Reasons" we support multiple astrological types.
- Monthly - Western - Greek Influenced Bablonian - https://en.wikipedia.org/wiki/Western_astrology

- Yearly - Chinese - https://en.wikipedia.org/wiki/Chinese_zodiac

All supported types have 100% code coverage because I dont want to ever have to think about this again.

# Please Note: I did not design GoLang.

Please Note: 100% code coverage in GoLang code is not a guarantee of a lack of bugs in GoLang code. This is in part due to GoLang's poor design mistakes when it comes to the "time" module.

One of the biggest issues many have with GoLang is that GoLang handles time formatting badly, using arbitrary format strings and date values that are, strictly speaking, not always compatiable with the worlds standards on issues of time.

This is a known issue that will create issues in the future for users of this library after I have been dead for a few thousand years... but as this is known issue, with known problems, I felt that simply getting this module to 100% code coverage was enough on my end as a consumer of these time api's to have done my due diligence as an engineer.

Yet as an engineer, I need to bubble up the risk, and communicate that trade off, summed up as follows, because GoLang has commited itself publicly to backward compatability:

Good: This library can work flawlessly for thousands of years using integer based constants, having been unit tested and fully proven in 100% of every possible tested case to be accurate and complete for that time.

Bad: After I have been dead for several thousand years, this library may still exist long enough that it would reach a time period that will create issues for its users due to solar clock drift and the overflowing nature of time itself. At that time, it might need a single line of code updated after a few thousand years of perfect utility, to mitgate a small defect in the golang time module.. unless they allready fixed it by then.

# How to Run The Unit Tests With coverage On Windows Using Standard GoLang

You may not have the coverage tools installed.  Install them (and upgrade golang) first.

```
go clean --testcache

go test -v ./... -coverprofile=coverage.out

go tool cover -html=coverage.out
```

Or on Windows, you can just run the "cover.bat" script by typeing "cover" in the module durectory.

You should be greeted with output that looks like this at the bottom of your console:

```
PASS
coverage: 100.0% of statements
ok      github.com/duaneking/gozodiacs  0.616s  coverage: 100.0% of statements
```

# The Code Coverage you did here seems to be the real big deal here, why do that?
You are right; The actual Zodiac problem is pretty simple and benign as it is mostly just some math on a constant table, and dealing with dates and numbers.

But code coverage by itself is an amazingly useful tool that actively works to help teams go faster when it's used correctly, so I wanted to create a simple example module that showed this and also presented some useful QA ideas.

The result is a simple and self contained package that limited external dependencies as much as possible for security and ease of education that I can just give a dev I work with a link to.
