# WorkSummary

**Building/running**

* Clone the repo
* Change directory into the cloned directory
* Run unit tests (optional): `go test -v ./...`
* Build: `go build .`
* Execute: `./WorkSummary input.json output.json`

**Next Steps**

If I had more time, here are some improvements I would make:
* Additional unit testing
  * Most of the important and tricky logic got tested, things like splitting shifts and looking for conflicts.  I'd like to have added a unit tests in additional places like for loading the data
* If this was truly for production, I would also want to add checks for things invalid or other date formats and other data errors
* Sort the shifts for then employee to reduce the conflict checking. Right now it is always linear, but with having them sorted the worst case would still be linear but the best case could be faster.
  * Given then time I had, I focused on a fully working solution where I could return to improve, if needed.
