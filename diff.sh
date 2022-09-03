#/!bin/bash
OLD_COUNT=$(wc -l < old_keys.txt)
NEW_COUNT=$(wc -l < new_keys.txt)

DIFF=$(($NEW_COUNT-$OLD_COUNT))

echo "There are $DIFF new keys"
