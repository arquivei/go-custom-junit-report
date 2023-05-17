# go-custom-junit-report

Create custom Junit XML reports in GO.

## What is this package?

go-custom-junit-report allows you to create a custom XML in the jUnit format.
This format is widely used by CI/CD tools like Gitlab and Bitbucket. While most
of the time the report generation is built into your test framework, you may
find yourself creating custom tests (like integration tests for example) and in
need of generating your own results.

## Supported jUnit format

**Example:**

```xml
<?xml version="1.0" encoding="UTF-8"?>
<testsuites id="testsuites#1" name="test_report" tests="4" failures="3" errors="3">
    <testsuite id="testsuite#1" name="suite 1" tests="2" failures="3" errors="3">
        <testcase id="case#1" name="case 1" classname="report.TestMakeReport"></testcase>
        <testcase id="case#2" name="case 2" classname="report.TestMakeReport">
            <failure message="msg1" type="type_fail">test failure 1</failure>
            <failure message="msg2" type="type_fail">test failure 2</failure>
            <failure message="msg3" type="type_fail">test failure 3</failure>
            <error message="msg4" type="type_err">test error 1</error>
            <error message="msg5" type="type_err">test error 2</error>
            <error message="msg6" type="type_err">test error 3</error>
        </testcase>
    </testsuite>
    <testsuite tests="2" failures="0" errors="0">
        <testcase></testcase>
        <testcase></testcase>
    </testsuite>
</testsuites>
```

**XML Elements:**

Test case elements are grouped into a test suite. Test suites elemnts are gouped
in the test suites element.

Test suites element:

| Name     | Description                | Optional | Observations       |
| ----     | -----------                | -------- | ------------       |
| id       | Test suites ID             | Yes      | Omitted when empty |
| name     | Test suites name           | Yes      | Omitted when empty |
| tests    | Total number of test cases | No       | Defaults to 0      |
| failures | Total number of failures   | No       | Defaults to 0      |
| errors   | Total number of errors     | No       | Defaults to 0      |
| time     | Total time                 | Yes      | Omitted when empty |

Test suite element:

| Name     | Description          | Optional | Observations       |
| ----     | -----------          | -------- | ------------       |
| id       | Test suite ID        | Yes      | Omitted when empty |
| name     | Test suite name      | Yes      | Omitted when empty |
| tests    | Number of test cases | No       | Defaults to 0      |
| failures | Number of failures   | No       | Defaults to 0      |
| errors   | Number of errors     | No       | Defaults to 0      |
| time     | Suite time           | Yes      | Omitted when empty |

Test case element:

| Name      | Description           | Optional | Observations       |
| ----      | -----------           | -------- | ------------       |
| id        | Test suite ID         | Yes      | Omitted when empty |
| name      | Test suite name       | Yes      | Omitted when empty |
| classname | Test suite class name | Yes      | Omitted when empty |
| failures  | Number of failures    | No       | Defaults to 0      |
| errors    | Number of errors      | No       | Defaults to 0      |
| time      | Test time             | Yes      | Omitted when empty |

Each test case element can contain the test case output.

Failure element:

| Name      | Description     | Optional | Observations       |
| ----      | -----------     | -------- | ------------       |
| message   | Failure message | Yes      | Omitted when empty |
| type      | Failure type    | Yes      | Omitted when empty |

The failure element can contain the failure output.

Error element:

| Name      | Description   | Optional | Observations       |
| ----      | -----------   | -------- | ------------       |
| message   | Error message | Yes      | Omitted when empty |
| type      | Error type    | Yes      | Omitted when empty |

The error element can contain the failure output.

## How to use it

```go
    suites := NewTestSuites("id", "name")
    suite := NewTestSuite("id", "name")
    suites.AddTestSuite(suite)

    testCase := NewTestCase("id", "name", "classname")

    // start tracking test duration
    testCase.Start()

    // test code
    ok, err := testCode()

    // finish tracking test duration
    testCase.End()

    if !ok {
        failure := NewFailure("failure message", "failure type", "failure output")
        testCase.AddFailure(failure)
    } else {
        // optional: set test output
        testCase.SetContent("output")
    }

    if err != nil {
        testErr := NewError("error message", "error type", "error output")
        testCase.AddError(testErr)
    }

    suite.AddTestCase(testCase)

    suites.SaveReport("filename.xml")
```