# It took me about 5 hours
romanValues = {
   "I":1,
   "V":5,
   "X":10,
   "L":50,
   "C":100,
   "D":500,
   "M":1000,
}

def calculate_values(values):
    individualValues = []

    for digit in values:
        individualValues.append(romanValues[digit])
    return individualValues


def set_sum_abst(arabicValues):
    currentIndex = 0
    for value in arabicValues:
        if currentIndex > 0 and (arabicValues[currentIndex-1] < value):
            arabicValues[currentIndex-1] = arabicValues[currentIndex-1] * -1
        currentIndex += 1
    return arabicValues

def convert_numbers(romanInput):
    allValues = calculate_values(romanInput)
    signedAllValues = set_sum_abst(allValues)
    totalSum = 0

    print(f'All values: {allValues}')
    
    for value in signedAllValues:
        totalSum = totalSum + value
    return totalSum

def is_valid_roman_number(text):
    response = False

    if is_syntactically_correct(text):
        if is_grammatically_correct(text):
            response = True
    else:
        response = False
    print(f'Is the number valid? {response}')

    return response

def is_grammatically_correct(text):
    threeDigitCounter = 0
    currentIndex = 0
    print(text)

    for x in text:
        if currentIndex > 0 and (text[currentIndex-1] == x):
            threeDigitCounter += 1
        else:
            threeDigitCounter = 0

        if currentIndex > 2 and (text[currentIndex-2] < x) and (text[currentIndex-1] < x) :
            print('No more than 3 consecutive identical digits as a subst')
            return False
        
        if threeDigitCounter == 3:
            print('No more than 3 consecutive identical digits are valid')
            return False
        
        currentIndex = currentIndex +1
    return True

def is_syntactically_correct(text):
    for x in text:
      if x not in romanValues :
        print('letra encontrada no valida: ', x)
        return False
    return True


if __name__ ==  '__main__':
    isValidNumber = False
    upperInput = ''

    while isValidNumber == False:

        inputText = input('Give a valid roman notation number: ')
        upperInput = inputText.upper()
        print('Input: ',upperInput)

        isValidNumber = is_valid_roman_number(upperInput)
      
    print(convert_numbers(upperInput))