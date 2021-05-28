package handler

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type NameResponse struct {
	Noun   string `json:"noun"`
	Name   string `json:"name"`
	Number string `json:"number"`
}

var (
	nouns []string = []string{"people", "history", "way", "art", "world", "information", "map", "two", "family", "government", "health", "system", "computer", "meat", "year", "thanks", "music", "person", "reading", "method", "data", "food", "understanding", "theory", "law", "bird", "literature", "problem", "software", "control", "knowledge", "power", "ability", "economics", "love", "internet", "television", "science", "library", "nature", "fact", "product", "idea", "temperature", "investment", "area", "society", "activity", "story", "industry", "media", "thing", "oven", "community", "definition", "safety", "quality", "development", "language", "management", "player", "variety", "video", "week", "security", "country", "exam", "movie", "organization", "equipment", "physics", "analysis", "policy", "series", "thought", "basis", "boyfriend", "direction", "strategy", "technology", "army", "camera", "freedom", "paper", "environment", "child", "instance", "month", "truth", "marketing", "university", "writing", "article", "department", "difference", "goal", "news", "audience", "fishing", "growth", "income", "marriage", "user", "combination", "failure", "meaning", "medicine", "philosophy", "teacher", "communication", "night", "chemistry", "disease", "disk", "energy", "nation", "road", "role", "soup", "advertising", "location", "success", "addition", "apartment", "education", "math", "moment", "painting", "politics", "attention", "decision", "event", "property", "shopping", "student", "wood", "competition", "distribution", "entertainment", "office", "population", "president", "unit", "category", "cigarette", "context", "introduction", "opportunity", "performance", "driver", "flight", "length", "magazine", "newspaper", "relationship", "teaching", "cell", "dealer", "finding", "lake", "member", "message", "phone", "scene", "appearance", "association", "concept", "customer", "death", "discussion", "housing", "inflation", "insurance", "mood", "woman", "advice", "blood", "effort", "expression", "importance", "opinion", "payment", "reality", "responsibility", "situation", "skill", "statement", "wealth", "application", "city", "county", "depth", "estate", "foundation", "grandmother", "heart", "perspective", "photo", "recipe", "studio", "topic", "collection", "depression", "imagination", "passion", "percentage", "resource", "setting", "ad", "agency", "college", "connection", "criticism", "debt", "description", "memory", "patience", "secretary", "solution", "administration", "aspect", "attitude", "director", "personality", "psychology", "recommendation", "response", "selection", "storage", "version", "alcohol", "argument", "complaint", "contract", "emphasis", "highway", "loss", "membership", "possession", "preparation", "steak", "union", "agreement", "cancer", "currency", "employment", "engineering", "entry", "interaction", "mixture", "preference", "region", "republic", "tradition", "virus", "actor", "classroom", "delivery", "device", "difficulty", "drama", "election", "engine", "football", "guidance", "hotel", "owner", "priority", "protection", "suggestion", "tension", "variation", "anxiety", "atmosphere", "awareness", "bath", "bread", "candidate", "climate", "comparison", "confusion", "construction", "elevator", "emotion", "employee", "employer", "guest", "height", "leadership", "mall", "manager", "operation", "recording", "sample", "transportation", "charity", "cousin", "disaster", "editor", "efficiency", "excitement", "extent", "feedback", "guitar", "homework", "leader", "mom", "outcome", "permission", "presentation", "promotion", "reflection", "refrigerator", "resolution", "revenue", "session", "singer", "tennis", "basket", "bonus", "cabinet", "childhood", "church", "clothes", "coffee", "dinner", "drawing", "hair", "hearing", "initiative", "judgment", "lab", "measurement", "mode", "mud", "orange", "poetry", "police", "possibility", "procedure", "queen", "ratio", "relation", "restaurant", "satisfaction", "sector", "signature", "significance", "song", "tooth", "town", "vehicle", "volume", "wife", "accident", "airport", "appointment", "arrival", "assumption", "baseball", "chapter", "committee", "conversation", "database", "enthusiasm", "error", "explanation", "farmer", "gate", "girl", "hall", "historian", "hospital", "injury", "instruction", "maintenance", "manufacturer", "meal", "perception", "pie", "poem", "presence", "proposal", "reception", "replacement", "revolution", "river", "son", "speech", "tea", "village", "warning", "winner", "worker", "writer", "assistance", "breath", "buyer", "chest", "chocolate", "conclusion", "contribution", "cookie", "courage", "dad", "desk", "drawer", "establishment", "examination", "garbage", "grocery", "honey", "impression", "improvement", "independence", "insect", "inspection", "inspector", "king", "ladder", "menu", "penalty", "piano", "potato", "profession", "professor", "quantity", "reaction", "requirement", "salad", "sister", "supermarket", "tongue", "weakness", "wedding", "affair", "ambition", "analyst", "apple", "assignment", "assistant", "bathroom", "bedroom", "beer", "birthday", "celebration", "championship", "cheek", "client", "consequence", "departure", "diamond", "dirt", "ear", "fortune", "friendship", "funeral", "gene", "girlfriend", "hat", "indication", "intention", "lady", "midnight", "negotiation", "obligation", "passenger", "pizza", "platform", "poet", "pollution", "recognition", "reputation", "shirt", "sir", "speaker", "stranger", "surgery", "sympathy", "tale", "throat", "trainer", "uncle", "youth", "time", "work", "film", "water", "money", "example", "while", "business", "study", "game", "life", "form", "air", "day", "place", "number", "part", "field", "fish", "back", "process", "heat", "hand", "experience", "job", "book", "end", "point", "type", "home", "economy", "value", "body", "market", "guide", "interest", "state", "radio", "course", "company", "price", "size", "card", "list", "mind", "trade", "line", "care", "group", "risk", "word", "fat", "force", "key", "light", "training", "name", "school", "top", "amount", "level", "order", "practice", "research", "sense", "service", "piece", "web", "boss", "sport", "fun", "house", "page", "term", "test", "answer", "sound", "focus", "matter", "kind", "soil", "board", "oil", "picture", "access", "garden", "range", "rate", "reason", "future", "site", "demand", "exercise", "image", "case", "cause", "coast", "action", "age", "bad", "boat", "record", "result", "section", "building", "mouse", "cash", "class", "nothing", "period", "plan", "store", "tax", "side", "subject", "space", "rule", "stock", "weather", "chance", "figure", "man", "model", "source", "beginning", "earth", "program", "chicken", "design", "feature", "head", "material", "purpose", "question", "rock", "salt", "act", "birth", "car", "dog", "object", "scale", "sun", "note", "profit", "rent", "speed", "style", "war", "bank", "craft", "half", "inside", "outside", "standard", "bus", "exchange", "eye", "fire", "position", "pressure", "stress", "advantage", "benefit", "box", "frame", "issue", "step", "cycle", "face", "item", "metal", "paint", "review", "room", "screen", "structure", "view", "account", "ball", "discipline", "medium", "share", "balance", "bit", "black", "bottom", "choice", "gift", "impact", "machine", "shape", "tool", "wind", "address", "average", "career", "culture", "morning", "pot", "sign", "table", "task", "condition", "contact", "credit", "egg", "hope", "ice", "network", "north", "square", "attempt", "date", "effect", "link", "post", "star", "voice", "capital", "challenge", "friend", "self", "shot", "brush", "couple", "debate", "exit", "front", "function", "lack", "living", "plant", "plastic", "spot", "summer", "taste", "theme", "track", "wing", "brain", "button", "click", "desire", "foot", "gas", "influence", "notice", "rain", "wall", "base", "damage", "distance", "feeling", "pair", "savings", "staff", "sugar", "target", "text", "animal", "author", "budget", "discount", "file", "ground", "lesson", "minute", "officer", "phase", "reference", "register", "sky", "stage", "stick", "title", "trouble", "bowl", "bridge", "campaign", "character", "club", "edge", "evidence", "fan", "letter", "lock", "maximum", "novel", "option", "pack", "park", "plenty", "quarter", "skin", "sort", "weight", "baby", "background", "carry", "dish", "factor", "fruit", "glass", "joint", "master", "muscle", "red", "strength", "traffic", "trip", "vegetable", "appeal", "chart", "gear", "ideal", "kitchen", "land", "log", "mother", "net", "party", "principle", "relative", "sale", "season", "signal", "spirit", "street", "tree", "wave", "belt", "bench", "commission", "copy", "drop", "minimum", "path", "progress", "project", "sea", "south", "status", "stuff", "ticket", "tour", "angle", "blue", "breakfast", "confidence", "daughter", "degree", "doctor", "dot", "dream", "duty", "essay", "father", "fee", "finance", "hour", "juice", "limit", "luck", "milk", "mouth", "peace", "pipe", "seat", "stable", "storm", "substance", "team", "trick", "afternoon", "bat", "beach", "blank", "catch", "chain", "consideration", "cream", "crew", "detail", "gold", "interview", "kid", "mark", "match", "mission", "pain", "pleasure", "score", "screw", "sex", "shop", "shower", "suit", "tone", "window", "agent", "band", "block", "bone", "calendar", "cap", "coat", "contest", "corner", "court", "cup", "district", "door", "east", "finger", "garage", "guarantee", "hole", "hook", "implement", "layer", "lecture", "lie", "manner", "meeting", "nose", "parking", "partner", "profile", "respect", "rice", "routine", "schedule", "swimming", "telephone", "tip", "winter", "airline", "bag", "battle", "bed", "bill", "bother", "cake", "code", "curve", "designer", "dimension", "dress", "ease", "emergency", "evening", "extension", "farm", "fight", "gap", "grade", "holiday", "horror", "horse", "host", "husband", "loan", "mistake", "mountain", "nail", "noise", "occasion", "package", "patient", "pause", "phrase", "proof", "race", "relief", "sand", "sentence", "shoulder", "smoke", "stomach", "string", "tourist", "towel", "vacation", "west", "wheel", "wine", "arm", "aside", "associate", "bet", "blow", "border", "branch", "breast", "brother", "buddy", "bunch", "chip", "coach", "cross", "document", "draft", "dust", "expert", "floor", "god", "golf", "habit", "iron", "judge", "knife", "landscape", "league", "mail", "mess", "native", "opening", "parent", "pattern", "pin", "pool", "pound", "request", "salary", "shame", "shelter", "shoe", "silver", "tackle", "tank", "trust", "assist", "bake", "bar", "bell", "bike", "blame", "boy", "brick", "chair", "closet", "clue", "collar", "comment", "conference", "devil", "diet", "fear", "fuel", "glove", "jacket", "lunch", "monitor", "mortgage", "nurse", "pace", "panic", "peak", "plane", "reward", "row", "sandwich", "shock", "spite", "spray", "surprise", "till", "transition", "weekend", "welcome", "yard", "alarm", "bend", "bicycle", "bite", "blind", "bottle", "cable", "candle", "clerk", "cloud", "concert", "counter", "flower", "grandfather", "harm", "knee", "lawyer", "leather", "load", "mirror", "neck", "pension", "plate", "purple", "ruin", "ship", "skirt", "slice", "snow", "specialist", "stroke", "switch", "trash", "tune", "zone", "anger", "award", "bid", "bitter", "boot", "bug", "camp", "candy", "carpet", "cat", "champion", "channel", "clock", "comfort", "cow", "crack", "engineer", "entrance", "fault", "grass", "guy", "hell", "highlight", "incident", "island", "joke", "jury", "leg", "lip", "mate", "motor", "nerve", "passage", "pen", "pride", "priest", "prize", "promise", "resident", "resort", "ring", "roof", "rope", "sail", "scheme", "script", "sock", "station", "toe", "tower", "truck", "witness", "a", "you", "it", "can", "will", "if", "one", "many", "most", "other", "use", "make", "good", "look", "help", "go", "great", "being", "few", "might", "still", "public", "read", "keep", "start", "give", "human", "local", "general", "she", "specific", "long", "play", "feel", "high", "tonight", "put", "common", "set", "change", "simple", "past", "big", "possible", "particular", "today", "major", "personal", "current", "national", "cut", "natural", "physical", "show", "try", "check", "second", "call", "move", "pay", "let", "increase", "single", "individual", "turn", "ask", "buy", "guard", "hold", "main", "offer", "potential", "professional", "international", "travel", "cook", "alternative", "following", "special", "working", "whole", "dance", "excuse", "cold", "commercial", "low", "purchase", "deal", "primary", "worth", "fall", "necessary", "positive", "produce", "search", "present", "spend", "talk", "creative", "tell", "cost", "drive", "green", "support", "glad", "remove", "return", "run", "complex", "due", "effective", "middle", "regular", "reserve", "independent", "leave", "original", "reach", "rest", "serve", "watch", "beautiful", "charge", "active", "break", "negative", "safe", "stay", "visit", "visual", "affect", "cover", "report", "rise", "walk", "white", "beyond", "junior", "pick", "unique", "anything", "classic", "final", "lift", "mix", "private", "stop", "teach", "western", "concern", "familiar", "fly", "official", "broad", "comfortable", "gain", "maybe", "rich", "save", "stand", "young", "fail", "heavy", "hello", "lead", "listen", "valuable", "worry", "handle", "leading", "meet", "release", "sell", "finish", "normal", "press", "ride", "secret", "spread", "spring", "tough", "wait", "brown", "deep", "display", "flow", "hit", "objective", "shoot", "touch", "cancel", "chemical", "cry", "dump", "extreme", "push", "conflict", "eat", "fill", "formal", "jump", "kick", "opposite", "pass", "pitch", "remote", "total", "treat", "vast", "abuse", "beat", "burn", "deposit", "print", "raise", "sleep", "somewhere", "advance", "anywhere", "consist", "dark", "double", "draw", "equal", "fix", "hire", "internal", "join", "kill", "sensitive", "tap", "win", "attack", "claim", "constant", "drag", "drink", "guess", "minor", "pull", "raw", "soft", "solid", "wear", "weird", "wonder", "annual", "count", "dead", "doubt", "feed", "forever", "impress", "nobody", "repeat", "round", "sing", "slide", "strip", "whereas", "wish", "combine", "command", "dig", "divide", "equivalent", "hang", "hunt", "initial", "march", "mention", "smell", "spiritual", "survey", "tie", "adult", "brief", "crazy", "escape", "gather", "hate", "prior", "repair", "rough", "sad", "scratch", "sick", "strike", "employ", "external", "hurt", "illegal", "laugh", "lay", "mobile", "nasty", "ordinary", "respond", "royal", "senior", "split", "strain", "struggle", "swim", "train", "upper", "wash", "yellow", "convert", "crash", "dependent", "fold", "funny", "grab", "hide", "miss", "permit", "quote", "recover", "resolve", "roll", "sink", "slip", "spare", "suspect", "sweet", "swing", "twist", "upstairs", "usual", "abroad", "brave", "calm", "concentrate", "estimate", "grand", "male", "mine", "prompt", "quiet", "refuse", "regret", "reveal", "rush", "shake", "shift", "shine", "steal", "suck", "surround", "anybody", "bear", "brilliant", "dare", "dear", "delay", "drunk", "female", "hurry", "inevitable", "invite", "kiss", "neat", "pop", "punch", "quit", "reply", "representative", "resist", "rip", "rub", "silly", "smile", "spell", "stretch", "stupid", "tear", "temporary", "tomorrow", "wake", "wrap", "yesterday"}

	names []string = []string{"liam", "noah", "oliver", "william", "elijah", "james", "benjamin", "lucas", "mason", "ethan", "alexander", "henry", "jacob", "michael", "daniel", "logan", "jackson", "sebastian", "jack", "aiden", "owen", "samuel", "matthew", "joseph", "levi", "mateo", "david", "john", "wyatt", "carter", "julian", "luke", "grayson", "isaac", "jayden", "theodore", "gabriel", "anthony", "dylan", "leo", "lincoln", "jaxon", "asher", "christopher", "josiah", "andrew", "thomas", "joshua", "ezra", "hudson", "charles", "caleb", "isaiah", "ryan", "nathan", "adrian", "christian", "maverick", "colton", "elias", "aaron", "eli", "landon", "jonathan", "nolan", "hunter", "cameron", "connor", "santiago", "jeremiah", "ezekiel", "angel", "roman", "easton", "miles", "robert", "jameson", "nicholas", "greyson", "cooper", "ian", "carson", "axel", "jaxson", "dominic", "leonardo", "luca", "austin", "jordan", "adam", "xavier", "jose", "jace", "everett", "declan", "evan", "kayden", "parker", "wesley", "kai", "brayden", "bryson", "weston", "jason", "emmett", "sawyer", "silas", "bennett", "brooks", "micah", "damian", "harrison", "waylon", "ayden", "vincent", "ryder", "kingston", "rowan", "george", "luis", "chase", "cole", "nathaniel", "zachary", "ashton", "braxton", "gavin", "tyler", "diego", "bentley", "amir", "beau", "gael", "carlos", "ryker", "jasper", "max", "juan", "ivan", "brandon", "jonah", "giovanni", "kaiden", "myles", "calvin", "lorenzo", "maxwell", "jayce", "kevin", "legend", "tristan", "jesus", "jude", "zion", "justin", "maddox", "abel", "king", "camden", "elliott", "malachi", "milo", "emmanuel", "karter", "rhett", "alex", "august", "river", "xander", "antonio", "brody", "finn", "elliot", "dean", "emiliano", "eric", "miguel", "arthur", "matteo", "graham", "alan", "nicolas", "blake", "thiago", "adriel", "victor", "joel", "timothy", "hayden", "judah", "abraham", "edward", "messiah", "zayden", "theo", "tucker", "grant", "richard", "alejandro", "steven", "jesse", "dawson", "bryce", "avery", "oscar", "patrick", "archer", "barrett", "leon", "colt", "charlie", "peter", "kaleb", "lukas", "beckett", "jeremy", "preston", "enzo", "luka", "andres", "marcus", "felix", "mark", "ace", "brantley", "atlas", "remington", "maximus", "matias", "walker", "kyrie", "griffin", "kenneth", "israel", "javier", "kyler", "jax", "amari", "zane", "emilio", "knox", "adonis", "aidan", "kaden", "paul", "omar", "brian", "louis", "caden", "maximiliano", "holden", "paxton", "nash", "bradley", "bryan", "simon", "phoenix", "lane", "josue", "colin", "rafael", "kyle", "riley", "jorge", "beckham", "cayden", "jaden", "emerson", "ronan", "karson", "arlo", "tobias", "brady", "clayton", "francisco", "zander", "erick", "walter", "daxton", "cash", "martin", "damien", "dallas", "cody", "chance", "jensen", "finley", "jett", "corbin", "kash", "reid", "kameron", "andre", "gunner", "jake", "hayes", "manuel", "prince", "bodhi", "cohen", "sean", "khalil", "hendrix", "derek", "cristian", "cruz", "kairo", "dante", "atticus", "killian", "stephen", "orion", "malakai", "ali", "eduardo", "fernando", "anderson", "angelo", "spencer", "gideon", "mario", "titus", "travis", "rylan", "kayson", "ricardo", "tanner", "malcolm", "raymond", "odin", "cesar", "lennox", "joaquin", "kane", "wade", "muhammad", "iker", "jaylen", "crew", "zayn", "hector", "ellis", "leonel", "cairo", "garrett", "romeo", "dakota", "edwin", "warren", "julius", "major", "donovan", "caiden", "tyson", "nico", "sergio", "nasir", "rory", "devin", "jaiden", "jared", "kason", "malik", "jeffrey", "ismael", "elian", "marshall", "lawson", "desmond", "winston", "nehemiah", "ari", "conner", "jay", "kade", "andy", "johnny", "jayceon", "marco", "seth", "ibrahim", "raiden", "collin", "edgar", "erik", "troy", "clark", "jaxton", "johnathan", "gregory", "russell", "royce", "fabian", "ezequiel", "noel", "pablo", "cade", "pedro", "sullivan", "trevor", "reed", "quinn", "frank", "harvey", "princeton", "zayne", "matthias", "conor", "sterling", "dax", "grady", "cyrus", "gage", "leland", "solomon", "emanuel", "niko", "ruben", "kasen", "mathias", "kashton", "franklin", "remy", "shane", "kendrick", "shawn", "otto", "armani", "keegan", "finnegan", "memphis", "bowen", "dominick", "kolton", "jamison", "allen", "philip", "tate", "peyton", "jase", "oakley", "rhys", "kyson", "adan", "esteban", "dalton", "gianni", "callum", "sage", "alexis", "milan", "moises", "jonas", "uriel", "colson", "marcos", "zaiden", "hank", "damon", "hugo", "ronin", "royal", "kamden", "dexter", "luciano", "alonzo", "augustus", "kamari", "eden", "roberto", "baker", "bruce", "kian", "albert", "frederick", "mohamed", "abram", "omari", "porter", "enrique", "alijah", "francis", "leonidas", "zachariah", "landen", "wilder", "apollo", "santino", "tatum", "pierce", "forrest", "corey", "derrick", "isaias", "kaison", "kieran", "arjun", "gunnar", "rocco", "emmitt", "abdiel", "braylen", "maximilian", "skyler", "phillip", "benson", "cannon", "deacon", "dorian", "asa", "moses", "ayaan", "jayson", "raul", "briggs", "armando", "nikolai", "cassius", "drew", "rodrigo", "raphael", "danny", "conrad", "moshe", "zyaire", "julio", "casey", "ronald", "scott", "callan", "roland", "saul", "jalen", "brycen", "ryland", "lawrence", "davis", "rowen", "zain", "ermias", "jaime", "duke", "stetson", "alec", "yusuf", "case", "trenton", "callen", "ariel", "jasiah", "soren", "dennis", "donald", "keith", "izaiah", "lewis", "kylan", "kobe", "makai", "rayan", "ford", "zaire", "landyn", "roy", "bo", "chris", "jamari", "ares", "mohammad", "darius", "drake", "tripp", "marcelo", "samson", "dustin", "layton", "gerardo", "johan", "kaysen", "keaton", "reece", "chandler", "lucca", "mack", "baylor", "kannon", "marvin", "huxley", "nixon", "tony", "cason", "mauricio", "quentin", "edison", "quincy", "ahmed", "finnley", "justice", "taylor", "gustavo", "brock", "ahmad", "kyree", "arturo", "nikolas", "boston", "sincere", "alessandro", "braylon", "colby", "leonard", "ridge", "trey", "aden", "leandro", "sam", "uriah", "ty", "sylas", "axton", "issac", "fletcher", "julien", "wells", "alden", "vihaan", "jamir", "valentino", "shepherd", "keanu", "hezekiah", "lionel", "kohen", "zaid", "alberto", "neil", "denver", "aarav", "brendan", "dillon", "koda", "sutton", "kingsley", "sonny", "alfredo", "wilson", "harry", "jaziel", "salvador", "cullen", "hamza", "dariel", "rex", "zeke", "mohammed", "nelson", "boone", "ricky", "santana", "cayson", "lance", "raylan", "lucian", "eliel", "alvin", "jagger", "braden", "curtis", "mathew", "jimmy", "kareem", "archie", "amos", "quinton", "yosef", "bodie", "jerry", "langston", "axl", "stanley", "clay", "douglas", "layne", "titan", "tomas", "houston", "darren", "lachlan", "kase", "korbin", "leighton", "joziah", "samir", "watson", "colten", "roger", "shiloh", "tommy", "mitchell", "azariah", "noe", "talon", "deandre", "lochlan", "joe", "carmelo", "otis", "randy", "byron", "chaim", "lennon", "devon", "nathanael", "bruno", "aryan", "flynn", "vicente", "brixton", "kyro", "brennan", "casen", "kenzo", "orlando", "castiel", "rayden", "ben", "grey", "jedidiah", "tadeo", "morgan", "augustine", "mekhi", "abdullah", "ramon", "saint", "emery", "maurice", "jefferson", "maximo", "koa", "ray", "jamie", "eddie", "guillermo", "onyx", "thaddeus", "wayne", "hassan", "alonso", "dash", "elisha", "jaxxon", "rohan", "carl", "kelvin", "jon", "larry", "reese", "aldo", "marcel", "melvin", "yousef", "aron", "kace", "vincenzo", "kellan", "miller", "jakob", "reign", "kellen", "kristopher", "ernesto", "briar", "gary", "trace", "joey", "clyde", "enoch", "jaxx", "crosby", "magnus", "fisher", "jadiel", "bronson", "eugene", "lee", "brecken", "atreus", "madden", "khari", "caspian", "ishaan", "kristian", "westley", "hugh", "kamryn", "musa", "rey", "thatcher", "alfred", "emory", "kye", "reyansh", "yahir", "cain", "mordechai", "zayd", "demetrius", "harley", "felipe", "louie", "branson", "graysen", "allan", "kole", "harold", "alvaro", "harlan", "amias", "brett", "khalid", "misael", "westin", "zechariah", "aydin", "kaiser", "lian", "bryant", "junior", "legacy", "ulises", "bellamy", "brayan", "kody", "ledger", "eliseo", "gordon", "london", "rocky", "valentin", "terry", "damari", "trent", "bentlee", "canaan", "gatlin", "kiaan", "franco", "eithan", "idris", "krew", "yehuda", "marlon", "rodney", "creed", "salvatore", "stefan", "tristen", "adrien", "jamal", "judson", "camilo", "kenny", "nova", "robin", "rudy", "van", "bjorn", "brodie", "mac", "jacoby", "sekani", "vivaan", "blaine", "ira", "ameer", "dominik", "alaric", "dane", "jeremias", "kyng", "reginald", "bobby", "kabir", "jairo", "alexzander", "benicio", "vance", "wallace", "zavier", "billy", "callahan", "dakari", "gerald", "turner", "bear", "jabari", "cory", "fox", "harlem", "jakari", "jeffery", "maxton", "ronnie", "yisroel", "zakai", "bridger", "remi", "arian", "blaze", "forest", "genesis", "jerome", "reuben", "wesson", "anders", "banks", "calum", "dayton", "kylen", "dangelo", "emir", "malakhi", "salem", "blaise", "tru", "boden", "kolten", "kylo", "aries", "henrik", "kalel", "landry", "marcellus", "zahir", "lyle", "dario", "rene", "terrance", "xzavier", "alfonso", "darian", "kylian", "maison", "foster", "keenan", "yahya", "heath", "javion", "jericho", "aziel", "darwin", "marquis", "mylo", "ambrose", "anakin", "jordy", "juelz", "toby", "yael", "azrael", "brentley", "tristian", "bode", "jovanni", "santos", "alistair", "braydon", "kamdyn", "marc", "mayson", "niklaus", "simeon", "colter", "davion", "leroy", "ayan", "dilan", "ephraim", "anson", "merrick", "wes", "will", "jaxen", "maxim", "howard", "jad", "jesiah", "ignacio", "zyon", "ahmir", "jair", "mustafa", "jermaine", "yadiel", "aayan", "dhruv", "seven", "stone", "rome"}
)

func Handler(w http.ResponseWriter, r *http.Request) {

	rand.Seed(time.Now().UnixNano())
	noun := strings.ToTitle(nouns[rand.Intn(len(nouns))])
	name := strings.ToTitle(names[rand.Intn(len(names))])
	number := strconv.Itoa(rand.Intn(99))

	resp := NameResponse{
		Noun:   noun,
		Name:   name,
		Number: number,
	}

	json.NewEncoder(w).Encode(resp)
}
