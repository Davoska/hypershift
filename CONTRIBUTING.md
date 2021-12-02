# How to contribute to Hypershift

Thanks for your interest in contributing to Hypershift. Here are some guidelines that help making the process more straight-forward for everyone:

* For small changes, you can just do the change and submit a pull request with it.
* For bigger changes (more than 200 lines of code diff), do not just do the change but ask for feedback on the idea and direction of the change first (Either in a GitHub issue or the #project-hypershift channel in the CoreOS slack). This avoids situations where big changes are submitted that are then declined or never reviewed, which is frustrating for everyone.
* Regardless of the size of the change, always explain how the change will improve the project.
* Make sure the "Why" and "How" are included in the message of each commit.
* Always keep refactorings (how we do something) separate from logic changes (what we do).
* Keep changes scoped to one thing and as minimal as possible. If you find additional things along the way that you feel should be improved, do that in a separate pull request. This helps ensure that you will get a timely review of your change, as a series of a small pull requests is a lot easier to review than one big pull request that changes 10 independent things for independent reasons.
* Before submitting the pull request on GitHub, look at your changes and try to view them from the eyes of a reviewer. Try to find the aspects that might not immediately make sense for someone else and explain them in the pull request description. Also ensure that the previous point (scoped to one thing and as minimal as possible) was followed.