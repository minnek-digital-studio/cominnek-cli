def question(question):
    i = 0
    while i < 2:
        answer = input(f"{question} (yes or no)")
        if any(answer.lower() == f for f in ["yes", 'y', '1', 'ye']):
            return True
        elif any(answer.lower() == f for f in ['no', 'n', '0']):
            return False
        else:
            i += 1
            if i < 10:
                print('Please enter yes or no')
            else:
                print("Nothing done")
                return False