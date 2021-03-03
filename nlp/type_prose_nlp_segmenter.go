package nlp

// Correspond to the Prose Document.Token part-of-speech tags.
// https://github.com/jdkato/prose#segmenting

// Catches most nouns, pronouns, verbs, adjectives, adverbs, and the determiner "not"
var ProseTagList_WordsOnly = []string{"RB", "DT", "NN", "NS", "NNP", "NNPS", "NNS", "JJ", "VB", "VBD", "VBG", "VBG", "VBN"}

var IgnoreMinifiedList = []string{"BE", "IS", "ISNT", "AR", "ARENT", "AINT", "HAVE", "AN", "A", "IF", "WILL"}

var ProseNouns = []string{"NN", "NS", "NNP", "NNPS", "NNS"}

var ProseVerbs = []string{"VB", "VBD", "VBG", "VBG", "VBN", "VBP", "VBZ"}

var ProseAdjectives = []string{"JJ", "JJR", "JJS"}

var ProseAdverbs = []string{"RB", "RBR", "RBS", "RP"}

var ProseDeterminers = []string{"DT"}

var ProseNounsVerbs = append(ProseNouns, ProseVerbs...)

var ProseNounsVerbsAdverbs = append(ProseNounsVerbs, ProseAdverbs...)

var ProseAdjectivesAdverbs = append(ProseAdjectives, ProseAdverbs...)

var ProseNounsVerbsAdjAdv = append(ProseNounsVerbs, ProseAdjectivesAdverbs...)

var ProseAdjAdvWithDeterminers = append(ProseAdjectivesAdverbs, ProseDeterminers...)
