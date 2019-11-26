package house

var subjects = []string{
  "house that Jack built",
  "malt",
  "rat",
  "cat",
  "dog",
  "cow with the crumpled horn",
  "maiden all forlorn",
  "man all tattered and torn",
  "priest all shaven and shorn",
  "rooster that crowed in the morn",
  "farmer sowing his corn",
  "horse and the hound and the horn",
}

var subject_actions = map[string]string{
  "malt": "lay in",
  "rat": "ate",
  "cat": "killed",
  "dog": "worried",
  "cow with the crumpled horn": "tossed",
  "maiden all forlorn": "milked",
  "man all tattered and torn": "kissed",
  "priest all shaven and shorn": "married",
  "rooster that crowed in the morn": "woke",
  "farmer sowing his corn": "kept",
  "horse and the hound and the horn": "belonged to",
}

func Verse(subject_index int) string {
  if subject_index == 1 {

    return "This is the house that Jack built."

  } else {

    current_subject_index := subject_index - 1
    current_subject := subjects[current_subject_index]
    current_subject_text := "This is the " + current_subject

    for current_subject_index >= 1 {

      current_subject = subjects[current_subject_index]
      current_subject_text += "\nthat " + subject_actions[current_subject] + " the " + subjects[current_subject_index - 1]
      current_subject_index--

    }
    return current_subject_text + "."
  }
}

func Song() string {
  
  song := Verse(1)
  
  for i:=2; i <= len(subjects) ; i++ {
    
    song = song + "\n\n" + Verse(i)
  }
  
  return song
}