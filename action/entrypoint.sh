#!/bin/sh

junit_directory="${1}"

for filename in "${junit_directory}"/*.xml; do
  junit2md -i "${filename}" -v=true -s=false -p=false                   >> "$GITHUB_STEP_SUMMARY";
done
echo ""                                                                 >> "$GITHUB_STEP_SUMMARY"
echo "## Summary of all tests"                                                                 >> "$GITHUB_STEP_SUMMARY"
echo ""                                                                 >> "$GITHUB_STEP_SUMMARY"
echo "<details>"                                                        >> "$GITHUB_STEP_SUMMARY"
echo "  <summary>Click here to see</summary>" >> "$GITHUB_STEP_SUMMARY"
echo ""                                                                 >> "$GITHUB_STEP_SUMMARY"
for filename in "${junit_directory}"/*.xml; do
  junit2md -i "${filename}" -s=false                                    >> "$GITHUB_STEP_SUMMARY";
done
echo "</details>"                                                       >> "$GITHUB_STEP_SUMMARY"
