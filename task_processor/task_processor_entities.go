package processor

/*

The goal of this file is to define what every task processor should have in common.
A task shouldn't be able to be executed by itself. So we need a task processor to handle
the execution of the task and the relaunching of the task if it fails.

*/
