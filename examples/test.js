import Shell from 'k6/x/shell';

export default function () {
  console.log("K6 shell extension enabled, version: " + Shell.version)
  const cmdcons = "ls"
  console.log(Shell.apply(cmdcons))
}
