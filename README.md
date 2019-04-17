# durations

some helpful time/duration funcs

This package is mostly sugar. If you want to do duration math with properly typed ints (think return values from other funcs) you have to cast that int to a time.Duration (int64). This package makes the invocations a
bit more fluent and hides the casts from view.
