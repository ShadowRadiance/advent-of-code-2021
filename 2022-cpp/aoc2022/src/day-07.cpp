#include <days.h>

#include <numeric>
#include <map>
#include <cassert>
#include <algorithm>
#include <iterator>

namespace day_07
{
  using std::string;
  using std::vector;
  using std::map;
  using std::to_string;
  using std::plus;

  using strings = vector<string>;

  class File
  {
  public:
    File(string& name, size_t size) : name_(name), size_(size) {}
    string name() { return name_; }
    size_t size() { return size_; }
  private:
    string name_;
    size_t size_;
  };

  class Dir
  {
  public:
    Dir(string name="", Dir* pParent = nullptr)
      : name_(name)
      , pParentDirectory_(pParent)
    {
    }

    size_t size() {
      return accumulate(
        files_.begin(), 
        files_.end(), 
        size_t{ 0 },
        [](size_t accum, auto& pair) { return accum + pair.second.size(); }
      );
    }

    size_t recursiveSize() {
      if (recursiveSizeCacheValid) return recursiveSizeCache;

      recursiveSizeCache = size() +
        accumulate(
          subDirectories_.begin(),
          subDirectories_.end(),
          size_t{ 0 },
          [](size_t accum, auto& pair) { return accum + pair.second.recursiveSize(); }
      );
      recursiveSizeCacheValid = true;
      return recursiveSizeCache;
    }
    // const _InIt _First, const _InIt _Last, _Ty _Val, _Fn _Reduce_op

    string name() { 
      return name_;
    }

    Dir* parent() {
      return pParentDirectory_;
    }

    Dir* changeToSubdirectory(string name) {
      assert(subDirectories_.contains(name));
      return &subDirectories_[name];
    }

    Dir* createSubdirectory(string name) {
      assert(!subDirectories_.contains(name));
      auto [itInserted, success] = 
        subDirectories_.insert({ name, {name, this} });
      return &(itInserted->second);
    }

    File* createFile(string name, size_t size) {
      assert(!files_.contains(name));
      invalidateRecursiveSizeCache();
      auto [itInserted, success] = 
        files_.insert({ name, { name, size } });
      return &(itInserted->second);
    }

    void invalidateRecursiveSizeCache() {
      recursiveSizeCacheValid = false;
      if (pParentDirectory_) {
        pParentDirectory_->invalidateRecursiveSizeCache();
      }
    }

  private:
    string name_;
    Dir* pParentDirectory_;
    map<string, Dir> subDirectories_;
    map<string, File> files_;

    size_t recursiveSizeCache{ 0 };
    bool recursiveSizeCacheValid{ true };
  };

  class Drive
  {
  public:
    Drive()
      : rootDirectory_(Dir{""})
      , currentDirectory_(&rootDirectory_)
    {}

    void changeToRoot() {
      currentDirectory_ = &rootDirectory_;
    }

    void changeToParent() {
      currentDirectory_ = currentDirectory_->parent();
    }

    void changeToSubdirectory(string name) {
      currentDirectory_ = currentDirectory_->changeToSubdirectory(name);
    }

    void createSubdirectory(string name) {
      Dir* newDir = currentDirectory_->createSubdirectory(name);
      allDirectories.push_back(newDir);
    }

    void createFile(string name, size_t size) {
      currentDirectory_->createFile(name, size);
    }

    vector<Dir*> findDirectoriesWithRecursiveSizeUnder(size_t maxSize) {
      vector<Dir*> result;

      copy_if(
        allDirectories.begin(),
        allDirectories.end(),
        back_inserter(result),
        [maxSize](Dir* pDir) { return pDir->recursiveSize() < maxSize; }
      );

      return result;
    }

    vector<Dir*> findDirectoriesWithRecursiveSizeOver(size_t minSize) {
      vector<Dir*> result;

      copy_if(
        allDirectories.begin(),
        allDirectories.end(),
        back_inserter(result),
        [minSize](Dir* pDir) { return pDir->recursiveSize() >= minSize; }
      );

      return result;
    }

    size_t size() {
      return rootDirectory_.recursiveSize();
    }
  private:
    Dir rootDirectory_;
    Dir* currentDirectory_;
    vector<Dir*> allDirectories;
  };

  class Executor
  {
  public:
    Executor(Drive& drive) : drive_(drive) {}

    void parseCommand(string s) {
      if (s.starts_with("$ cd")) {
        size_t last_space_index = s.find_last_of(' ');
        string last_word = s.substr(last_space_index + 1);

        if (last_word == "/") {
          drive_.changeToRoot();
          return;
        }
        if (last_word == "..") {
          drive_.changeToParent();
          return;
        }
        drive_.changeToSubdirectory(last_word);
        return;
      }

      if (s == "$ ls" || s == "") {
        return;
      }

      if (s.starts_with("dir")) {
        size_t last_space_index = s.find_last_of(' ');
        string last_word = s.substr(last_space_index + 1);

        drive_.createSubdirectory(last_word);
        return;
      }

      size_t last_space_index = s.find_last_of(' ');
      string first_word = s.substr(0, last_space_index);
      size_t filesize = stoul(first_word);
      string last_word = s.substr(last_space_index + 1);
      drive_.createFile(last_word, filesize);
    }
  private:
    Drive& drive_;
  };

  void buildDrive(Drive& drive, const strings& input_data) {
    Executor executor(drive);

    for (const string& s : input_data) {
      executor.parseCommand(s);
    }
  }

  string answer_a(const strings& input_data)
  {
    Drive drive;
    buildDrive(drive, input_data);

    // To begin, find all of the directories with a total size of at 
    // most 100000, then calculate the sum of their total sizes. 
    // In the example above, these directories are a and e; 
    // the sum of their total sizes is 95437 (94853 + 584). 
    // (As in this example, this process can count files more than once!)
    
    vector<Dir*> directories = 
      drive.findDirectoriesWithRecursiveSizeUnder(100000);

    size_t sum = accumulate(
      directories.begin(),
      directories.end(),
      size_t{ 0 },
      [](size_t accum, Dir* pDir) {
        return accum + pDir->recursiveSize();
      }
    );

    return to_string(sum);
  }

  string answer_b(const vector<string>& input_data)
  {
    Drive drive;
    buildDrive(drive, input_data);

    // The total disk space available to the filesystem is 70_000_000.
    // To run the update, you need unused space of at least 30_000_000. 
    // You need to find a directory you can delete that will free up 
    // enough space to run the update.

    const size_t DRIVE_MAX{ 70000000UL };
    const size_t REQUIRED_FREE{ 30000000UL };

    size_t currentUsage = drive.size();
    size_t currentFree = DRIVE_MAX - currentUsage;
    size_t requiredToFreeUp = REQUIRED_FREE - currentFree;

    vector<Dir*> potentialDirectories =
      drive.findDirectoriesWithRecursiveSizeOver(requiredToFreeUp);

    // Find the smallest directory that, if deleted, would free up 
    // enough space on the filesystem to run the update. What is the 
    // total size of that directory?
    auto it = min_element(
      potentialDirectories.begin(),
      potentialDirectories.end(),
      [](Dir* a, Dir* b) { return a->recursiveSize() < b->recursiveSize(); }
    );

    if (it == potentialDirectories.end()) {
      return "COULDN'T FIND A DIRECTORY LARGE ENOUGH";
    }

    Dir* smallest = *it;
    return to_string(smallest->recursiveSize());
  }
}
