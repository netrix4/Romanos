# It took me about 5 hours XL -> 40
#
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

def convert_units(strToConvert):
    digit = int(strToConvert)
    value = ""
    
    if digit < 4:
        for i in range(digit):
            value = value + "I"
    if digit == 4:
        value = 'IV'
    if digit > 4:
        value = 'V'
        if digit > 5:
            if digit == 9:
                value = "IX"
            else:
                for x in range(digit-5):
                  value = value + "I"

    return value

def convert_tens(strToConvert):
    digit = int(strToConvert)
    value = ""

    if digit < 4:
        for i in range(digit):
            value += "X"
    if digit == 4:
        value = "XL"
    if digit > 4:
        value = 'L'
        if digit > 5:
            if digit == 9:
                value = "XC"
            else:
                for x in range(digit-5):
                  value = value + "X"
    return value

def convert_hundreds(strToConvert):
    digit = int(strToConvert)
    value = ""

    if digit < 4:
        for i in range(digit):
            value += "C"
    if digit == 4:
        value = "CD"
    if digit > 4:
        value = 'D'
        if digit > 5:
            if digit == 9:
                value = "CM"
            else:
                for x in range(digit-5):
                  value = value + "C"
    return value

def convert_thousands(strToConvert):
    digit = int(strToConvert)
    value = ""

    for x in range(digit):
        value += value + "M"
    return value

def convert_arabig_to_roman(numberInText):
    totalRomanNumber = ""
    newNumber = ""

    digitCounter = 0

    for number in reversed(numberInText):
        digitCounter +=1

        if digitCounter == 1:
            newNumber = convert_units(number)
        if digitCounter == 2:
            newNumber = convert_tens(number)
        if digitCounter == 3:
            newNumber = convert_hundreds(number)
        if digitCounter == 4:
            newNumber = convert_thousands(number)
        totalRomanNumber = newNumber + totalRomanNumber

    return totalRomanNumber

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

def is_valid_arabig_number(numberStr):
    try:
      inputTextNumber = int(numberStr)
      return True
    except:
      return False

def is_grammatically_correct(text):
    threeDigitCounter = 0
    currentIndex = 0
    print(text)

    for x in text:
        if currentIndex > 0 and (text[currentIndex-1] == x):
            threeDigitCounter += 1
        else:
            threeDigitCounter = 0

        # if currentIndex > 2 and (text[currentIndex-2] < x) and (text[currentIndex-1] < x) :
        #     print('No more than 3 consecutive identical digits as a subst')
        #     return False
        
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

def roman_2_arabic():
    isValidNumber = False
    while isValidNumber == False:

        inputText = input('Give a valid roman notation number: ')
        upperInput = inputText.upper()
        print('Input: ',upperInput)
        isValidNumber = is_valid_roman_number(upperInput)
      
    result = convert_numbers(upperInput)
    print(f'Your number is: {result}')

#///Nuevo
def arabic_2_roman():
    isValidNumber = False

    while isValidNumber == False:
      inputTextNumber = input('Give a valid arabig notation number: ')
      isValidNumber = is_valid_arabig_number(inputTextNumber)

    result = convert_arabig_to_roman(inputTextNumber)
    print(f'Your number is: {result}')

if __name__ ==  '__main__':
    upperInput = ''
    convertionOption = input('(1) Roman to arabic: MCMXCVII -> 1997\n(2) Arabic to roman: 1997 -> MCMXCVII\nSelect an option: ')

    if convertionOption == '1':
      roman_2_arabic()
    if convertionOption == '2':
      arabic_2_roman()
