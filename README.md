# Implementation of the parallel sort algorithm in Go 

## Common
The application implements parallel sorting of 1000 names from the input text file.

## Algorithm

Conceptually, a merge sort works as follows:

- Divide the unsorted list into n sublists, each containing one element (a list of one element is considered sorted).
- Repeatedly merge sublists to produce new sorted sublists until there is only one sublist remaining. This will be the sorted list.