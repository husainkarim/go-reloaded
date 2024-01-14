# GO-RELOADED

## Project Overview

This project involves building a tool that automates text modifications for improved clarity and consistency. It will utilize functions from previous repositories and undergo peer-review for quality assurance.

## Key Features

- Hexadecimal to Decimal Conversion: Replaces hexadecimal numbers with their decimal equivalents.
- Binary to Decimal Conversion: Converts binary numbers to decimal.
Case Manipulation:
- Uppercase (up): Converts words to uppercase.
- Lowercase (low): Converts words to lowercase.
- Capitalize (cap): Capitalizes the first letter of words.
- Supports multiple word modifications with (up, "number"), (low, "number"), and (cap, "number").
- Punctuation Formatting:
    - Adjusts spacing for punctuation marks (., ,, !, ?, :, ; ).
    - Handles multiple punctuation groups (e.g., ...) appropriately.
    - Correctly places single quotes (') around words.
- Article Correction: Changes "a" to "an" when followed by a vowel or "h."

## Usage

1. Provide Input and Output Files:

```
go run . sample.txt result.txt
```

2. Review Modified Output: The tool will apply the specified modifications and save the results to the specified output file.

## Testing

- Create comprehensive tests: Include varied scenarios to ensure functionality.
- Utilize tests for peer-review: Share tests for project evaluation and feedback.

## Collaboration and Feedback

We welcome contributions and suggestions for improvement! Share your feedback to enhance this tool's effectiveness.
