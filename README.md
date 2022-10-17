## To reproduce

```bash
./pants test src/example:
```

Output:
```raw
13:21:52.60 [INFO] Initializing scheduler...
13:21:52.66 [INFO] Scheduler initialized.
13:21:53.13 [ERROR] 1 Exception encountered:

  Exception: The dependency graph contained a cycle:

  @rule(pants.backend.go.util_rules.build_pkg_target.setup_build_go_package_target_request(src/example:example)) <-
  @rule(pants.backend.go.util_rules.build_pkg_target.setup_build_go_package_target_request(src/example/testutils:testutils))
  @rule(pants.backend.go.util_rules.build_pkg_target.setup_build_go_package_target_request(src/example:example)) <-

If the dependencies in the above path are for your BUILD targets, you may need to use more granular targets or replace BUILD target dependencies with file dependencies. If they are not for your BUILD targets, then please file a Github issue!

See https://www.pantsbuild.org/v2.13/docs/targets#dependencies-and-dependency-inference for more information.
```
