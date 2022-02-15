math.randomseed(os.time())
num = math.random(1,100)
guessed = false
 
 
while (guessed == false)
do
    guess = io.read("*n")
    
    if guess == num then
        print("Correct.\n")
    elseif guess > num then
        print("Too high.\n")
    else
        print("Too low.\n")
    end
end
