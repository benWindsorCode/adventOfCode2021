processed ← ⍎¨ ⊃⎕NGET 'input.txt'1
diff←{(¯1↓1⌽⍵)-(¯1↓⍵)}
answer ← +/ 0 >⍨ diff processed