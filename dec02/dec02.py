games = [line.strip() for line in open('input2')]
score1 = 0
score2 = 0
for game in games:
    them,me = game.split()
    print(them, me)
    score1 += {'X': 1, 'Y': 2, 'Z': 3}[me]
    score1 += {('A', 'X'): 3, ('A', 'Y'): 6, ('A', 'Z'): 0,
              ('B', 'X'): 0, ('B', 'Y'): 3, ('B', 'Z'): 6,
              ('C', 'X'): 6, ('C', 'Y'): 0, ('C', 'Z'): 3,
              }[(them, me)]
    
    score2 += {'X': 0, 'Y': 3, 'Z': 6}[me]
    score2 += {('A', 'X'): 3, ('A', 'Y'): 1, ('A', 'Z'): 2,
              ('B', 'X'): 1, ('B', 'Y'): 2, ('B', 'Z'): 3,
              ('C', 'X'): 2, ('C', 'Y'): 3, ('C', 'Z'): 1,
              }[(them, me)]
    
    print(score1)
    print(score2)