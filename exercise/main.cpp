/**
 * 自动图书馆 主程序
 */

#pragma warning(disable:4786)
#include <iostream>
#include <cstring>
#include <string>
#include <fstream>
#include <map>
#include <typeinfo> 
#include<time.h>

const int LIMIT = 3;

enum State { READER, LIB, SHELF };

class Object{
};

using namespace std;

template <class T>
class Queue  
{
private:
	int front;
	int rear;
	int size;
	T * data;

public:
	Queue(int s = 100); 
	virtual ~Queue();

	void insert(const T & x);
	T remove();
	int is_empty() const;
	int is_full() const;
};

template <class T>
Queue<T>::Queue(int s): size(s+1), front(0), rear(0)
{
	data = new T[size];
}

template <class T>
Queue<T>::~Queue()
{
	delete [] data;
}

template <class T>
void Queue<T>::insert(const T & x)
{
	data[rear ++ % size] = x;
}

template <class T>
T Queue<T>::remove()
{
	return data[front ++ % size];
}

template <class T>
int Queue<T>::is_empty() const
{
	return front == rear;
}

template <class T>
int Queue<T>::is_full() const
{
	return ((rear + 1) % size == front);
}

class BookData : public Object
{
private:
	long bookID;
	char name[40];
	State state;
	long PIN;

public:
	BookData(long bookID, const char * name, State state = SHELF, long PIN = 0)
	{
		SetID(bookID);
		SetName(name);
		SetState(state);
		SetPIN(PIN);
	}

	BookData()
	{
		bookID = 0;
		PIN = 0;
	}

	long GetID() const
	{
		return bookID;
	}

	const char * GetName() const
	{
		return name;
	}

	State GetState() const
	{
		return state;
	}

	long GetPIN() const
	{
		return PIN;
	}

	void ShowData()
	{
		cout << bookID << "\t" << name << "\t" << state << "\t" << PIN << endl;
	}

	void SetID(long bookID)
	{
		this->bookID = bookID;
	}

	void SetName(const char * name)
	{
		strcpy(this->name, name);
	}

	void SetState(State state)
	{
		this->state = state;
	}

	void SetPIN(long PIN)
	{
		this->PIN = PIN;
	}
};

template <class T>
class Database : public Object
{
private:
	fstream file;
	char filename[40];
	long fileLen, recSize;
	typedef map< long, T, less<long> > mtype;
	mtype recMap;

public:
	Database(const char * filename)
	{
		strcpy(this->filename, filename);
		file.open(filename, ios::in | ios::binary);
		recSize = sizeof(T);
		if (file.is_open())
		{
			file.seekg(0, ios::end);
			if ((fileLen = file.tellg()) > 0)
			{
				T obj;
				file.seekg(0, ios::beg);
				do
				{
					file.read((char *)&obj, recSize);
					recMap.insert(typename mtype::value_type(obj.GetID(), obj));
				} while (file.tellg() < fileLen);
			}

			file.close();
		}
	}

	void SaveMap()
	{
		typename mtype::const_iterator iter;
		
		file.open(filename, ios::out | ios::binary | ios:: trunc);
		for (iter = recMap.begin(); iter != recMap.end(); ++iter)
		{
			file.write((char *)&iter->second, recSize);
		}

		file.close();
	}

	void Insert(const T & obj)
	{
		recMap.insert(typename mtype::value_type(obj.GetID(), obj));
	}

	T * Query(long objID)
	{
		typename mtype::iterator iter;
		iter = recMap.find(objID);

		if (iter == recMap.end())
		{
			cout << objID << " not found!" << endl;
			return NULL;
		} else
		{
			return &(iter->second);
		}
	}

	T * QueryName(const char * objName)
	{
		typename mtype::iterator iter;
		for (iter = recMap.begin(); iter != recMap.end(); iter ++)
		{
			if (strstr((iter->second).GetName(), objName) != NULL)
			{
				return &(iter->second);
			}
		}

		cout << objName << " in " << typeid(T).name() << " not found " << endl;
		return NULL;
	}

	void Delete(long objID)
	{
		Query(objID);
		recMap.erase(objID);
	}

	void ShowAllData()
	{
		typename mtype::iterator iter;
		T obj;

		cout << "Data in " << typeid(T).name() << ":" << endl;
		for (iter = recMap.begin(); iter != recMap.end(); iter ++)
		{
			(iter->second).ShowData();
		}
	}
};

class Date : public Object
{
	int year, month, day;

	int DayOfMonth (int y, int m) const
	{
		int d = 0;
		switch (m)
		{
		case 1: 
		case 3:
		case 5:
		case 7: 
		case 8:
		case 10:
		case 12: d = 31;  break;
		case 4: 
		case 6:
		case 9:
		case 11: d = 30;  break;
		case 2: d = 28 + IsLeapYear(y);  break;
		}

		return d;
	}

public:
	Date()
	{
		time_t curTime = time(NULL);
		tm time = *localtime(&curTime);
		day = time.tm_mday;
		month = time.tm_mon + 1;
		year = time.tm_year + 1900;
	}

	Date(int y, int m, int d) : year(y), month(m), day(d)
	{
		if ((y <= 0 ) || (m <= 0) || m > 12 || d <= 0 || d > DayOfMonth(y, m))
		{
			year = 1900;
			month = day = 1;
		}
	}

	virtual ~Date()
	{
	}

	int GetYear() const
	{
		return year;
	}

	int GetMonth() const
	{
		return month;
	}

	int GetDay() const
	{
		return day;
	}

	bool IsLeapYear() const
	{
		return IsLeapYear(year);
	}

	bool IsLeapYear(const int y) const
	{
		return year % 400 ? (year % 100 ? (year % 4 ? false : true) : false): true;
	}

	void Display() const
	{
		cout << year << "-" << month << "-" << day << endl;
	}
};

class Reader : public Object
{
private:
	long PIN;
	long bookID[LIMIT];
	int num;

public:
	Reader(long PIN = 0, const int num = 0): PIN(PIN), num(num)
	{
	}

	int AddBook(const long bookID)
	{
		if (num < LIMIT)
		{
			this->bookID[num] = bookID;
			cout << "Book " << bookID << " added!" << endl;
			num ++;
			return 1;
		} else
		{
			cout << "Cart is Full" << endl;
			return 0;
		}
	}

	long CheckOut()
	{
		return bookID[--num];
	}

	void ShowCart()
	{
		for (int i = 0; i < num; ++i)
		{
			cout << bookID[i] << endl;
		}
	}

	int GetNum() const
	{
		return num;
	}
};


class ReaderData : public Object
{
protected:
	long PIN;
	char name[20];
	long bookList[LIMIT];
	Date borrowDate[LIMIT];
	int num;

public:
	ReaderData(int PIN, const char * name)
	{
		SetID(PIN);
		SetName(name);
		num = 0;
		for (int i = 0; i < LIMIT; i ++)
			bookList[i] = 0;
	}

	ReaderData()
	{
		PIN = 0; num = 0;
		for (int i = 0; i < LIMIT; i ++)
			bookList[i] = 0;
	}

	void SetID(long PIN)
	{
		this->PIN = PIN;
	}

	void SetName(const char * name)
	{
		strcpy(this->name, name);
	}

	long GetID() const
	{
		return PIN;
	}

	const char * GetName() const
	{
		return name;
	}

	int GetNum() const
	{
		return num;
	}

	long ReturnBook(long bookID)
	{
		for (int i = 0; i < LIMIT; i ++)
		{
			if (bookList[i] == bookID)
			{
				bookList[i] = 0; 
				-- num;
				return bookID;
			}
		}

		return 0;
	}

	long BorrowBook(long bookID)
	{
		for (int i = 0; i < LIMIT; i ++)
		{
			if (bookList[i] == 0)
			{
				bookList[i] = bookID;
				borrowDate[i] = Date();
				num ++;
				return bookID;
			}
		}

		cout << "Book reach the limit!" << endl;
		return 0;
	}

	void ShowData()
	{
		cout << PIN << '\t' << name << endl;
		for (int i = 0; i < LIMIT; i ++)
		{
			if (bookList[i])
			{
				cout << i + 1 << ":" << bookList[i] << '\t';
				borrowDate[i].Display();
			}
		}
	}
};

class Librarian: public Object
{
private:
	long ID;
	char name[20];
	char passwd[9];

public:
	Librarian(long ID, const char * name) : ID(ID)
	{
		strcpy(this->name, name);
		strcpy(passwd, "abc");
	}

	char login()
	{
		char pw[9];
		for (int i = 0; i < 3; i ++)
		{
			cout << "Enter Password:";
			cin >> pw;
			if (strcmp(pw, passwd) == 0)
				return 'X';
		}
		cout << "Login Failed!" << endl;
		return 'E';
	}
};

// 动作相关类
class Action {
public:
	virtual void execute() = 0;
};

//应用程序类
class Application {
private:
	Queue<Action*> actions;  // 定义一个Action队列，应用程序会不停的从队列里取Action并执行
public:
	Application(Action * pAction = NULL) : actions(10000) {
		if (NULL != pAction)
		    actions.insert(pAction);
	}
	virtual ~Application() {
	}

	virtual void run() {
		while(!actions.is_empty()) {
			Action * pAction = actions.remove();
			if (NULL != pAction)
			    pAction->execute();
		}
	}

	virtual void pushAction(Action * pAction) {
		actions.insert(pAction);
	}
};

// 菜单命令回调函数指针
typedef void (Application::*ActionProcess)();

// 菜单和菜单项
class MenuItem {
private:
	char  text[64];
	ActionProcess callback;
	Application * app;

public:
	MenuItem(char * text = "") {
		strcpy(this->text, text);
		callback = NULL;
	}

	void setData(char * text, Application * app, ActionProcess callback) {
		strcpy(this->text, text);
		this->app = app;
		this->callback = callback;
	}
	char * getText() {
		return text;
	}

	void click() {
		if (NULL != callback) {
			(app->*callback)();
		}
	}
};

class Menu : public Action {
private:
	char  menuName[12];
	MenuItem * items;
	int num;
public:
	Menu(char * menuName,  char * texts[], Application * app, ActionProcess * callbacks, int num) {
		strcpy(this->menuName, menuName);
		this->num = num;
		items = new MenuItem[num];
		for (int i = 0; i < num; i ++) {
			items[i].setData(texts[i], app, callbacks[i]);
		}
	}

	virtual void showMenu() {
		printf("----------------%s----------------\n", menuName);
		for (int i = 0; i < num; i ++) {
			printf("%d:%s\n", i, items[i].getText());
		}
	}

	virtual void listen() {
		int index = -1;
		
		do {
			cout << "\n请选择菜单项:";
			if (cin >> index) {
				if (index >= 0 && index < num) {
					items[index].click();
					break;
				}
			} else {
				cin.clear();             // 错误标志位清除
				cin.ignore(9999, '\n');  // 忽略缓冲区数据
			}
		} while (true);
	}

	virtual void execute() {
		this->showMenu();
		this->listen();
	}

	virtual ~Menu() {
		if (NULL != items) {
			delete[] items;
			items = NULL;
		}
	}
};

// 自定义应用程序
class LibApp : public Application {
private:
	Database<BookData>    bookBase;
	Database<ReaderData>  readerBase;

	Menu * pMainMenu;
	Menu * pLibMenu;
	Menu * pReaderMenu;

	Librarian * mgr;
	Reader * rdr;
	long PIN;

public:
	// 构造函数
	LibApp() : bookBase("./books.dat"), readerBase("./readers.dat") {
		char * mainMenuTexts[]           = {"Librarian entry", "Reader entry", "Return book", "Exit"};
		ActionProcess mainMenuCallback[] = {(ActionProcess)&LibApp::mainLibEntry, (ActionProcess)&LibApp::mainReaderEntry, (ActionProcess)&LibApp::mainReturnBook, (ActionProcess)&LibApp::mainExit};
		pMainMenu = new Menu("主菜单", mainMenuTexts, this, mainMenuCallback, 4);

		char * libMenuTexts[] = {"Add reader", "Add Book", "Query Reader", "Put book to shelf", "Show All", "Exit"};
		ActionProcess libMenuCallback[] = {(ActionProcess)&LibApp::libAddReader, (ActionProcess)&LibApp::libAddBook, (ActionProcess)&LibApp::libQueryReader, (ActionProcess)&LibApp::libPutBook, (ActionProcess)&LibApp::libShowAll, (ActionProcess)&LibApp::libExit};
		pLibMenu = new Menu("管理员菜单", libMenuTexts, this, libMenuCallback, 6);

		char * readerMenuTexts[] = {"Add book to Cart", "Check Out", "Query Book By Name", "List My Books", "Show My Cart", "Exit"};
		ActionProcess readerMenuCallback[] = {(ActionProcess)&LibApp::readerAddBook, (ActionProcess)&LibApp::readerCheckOut, (ActionProcess)&LibApp::readerQueryBook, (ActionProcess)&LibApp::readerListBook,(ActionProcess) &LibApp::readerCart, (ActionProcess)&LibApp::readerExit};
		pReaderMenu = new Menu("读者菜单", readerMenuTexts, this, readerMenuCallback, 6);

		mgr = NULL;
		rdr = NULL;

		// 初始化命令
		pushAction(pMainMenu);
	}

	~LibApp() {
		if (NULL != pMainMenu) {
		    delete pMainMenu;
			pMainMenu = NULL;
		}
		if (NULL != pLibMenu) {
		    delete pLibMenu;
			pLibMenu = NULL;
		}
		if (NULL != pReaderMenu) {
		    delete pReaderMenu;
			pReaderMenu = NULL;
		}
		if (NULL != mgr) {
		    delete mgr;
			mgr = NULL;
		}
		if (NULL != rdr) {
		    delete rdr;
			rdr = NULL;
		}
	}

	// 菜单处理函数
	//main
	void mainLibEntry() {
		mgr = new Librarian(101, "yjc");
		char ret = mgr->login(); 
		if (ret == 'X') {
			pushAction(pLibMenu);
		} else {
			cout << "您的密码三次输入错误！" << endl;
			pushAction(pMainMenu);
		}
	}
	void mainReaderEntry() {
		int i;

		for (i = 0; i < 3; i ++) {
			cout << "Input PIN:";
			cin >> PIN;
			if (readerBase.Query(PIN) != NULL) {
				rdr = new Reader(PIN);
				break;
			}
		}
		if (i == 3) {
			cout << "Check In Failed!" << endl;
			pushAction(pMainMenu);
		} else {
			cout << "读者菜单：" << endl;
			pushAction(pReaderMenu);
		}
	}
	void mainReturnBook() {
		long PIN, bookID;
		cout << "Input a book ID:";
		cin >> bookID;

		if (bookBase.Query(bookID) != NULL) {
			bookBase.Query(bookID)->SetState(LIB);
			if ((PIN = bookBase.Query(bookID)->GetPIN()) > 0) {
				readerBase.Query(PIN)->ReturnBook(bookID);
			}
		}
		pushAction(pMainMenu);
	}
	void mainExit() {
		cout << "程序结束" << endl;
		exit(0);
	}
	//lib
	void libAddReader() {
		long ID;
		char name[40];

		cout << "Give a reader PIN and input a name:";
		cin >> ID;
		cin.ignore();
		cin.get(name, 20, '\n');
		readerBase.Insert(ReaderData(ID, name));

		pushAction(pLibMenu);
	}
	void libAddBook() {
		long ID;
		char name[40];

		cout << "Input a book ID and name:";
		cin >> ID;
		cin.ignore();
		cin.get(name, 40, '\n');
		bookBase.Insert(BookData(ID, name));

		pushAction(pLibMenu);
	}
	void libQueryReader() {
		long ID;

		cout << "Input a reader's PIN:";
		cin >> ID;
		if (readerBase.Query(ID) == NULL)
		{
			cout << "No such a reader! " << endl;
		} else
		{
			readerBase.Query(ID)->ShowData();
		}

		pushAction(pLibMenu);
	}
	void libPutBook() {
		long ID;

		cout << "Input a book ID:";
		cin >> ID;
		if (bookBase.Query(ID) == NULL)
		{
			cout << "No such a book!" << endl;
		} else
		{
			bookBase.Query(ID)->SetState(SHELF);
		}

		pushAction(pLibMenu);
	}
	void libShowAll() {
		readerBase.ShowAllData();
		bookBase.ShowAllData();

		pushAction(pLibMenu);
	}
	void libExit() {
		pushAction(pMainMenu);
	}
	//reader
	void readerAddBook() {
		long bookID;

		cout << "Input a book ID:";
		cin >> bookID;
		if (bookBase.Query(bookID) != NULL && bookBase.Query(bookID)->GetState() == SHELF)
		{
			if (rdr->AddBook(bookID))
				bookBase.Query(bookID)->SetState(LIB);
		}

		pushAction(pReaderMenu);
	}
	void readerCheckOut() {
		int i, t1, t2;
		long bookID;

		t1 = rdr->GetNum();
		t2 = readerBase.Query(PIN)->GetNum();

		if (t1 > 0 && t2 < LIMIT)
		{
			cout << PIN << " " << readerBase.Query(PIN)->GetName() << " Book List" << endl;
			for (i = 0; i < t1 && i < (LIMIT - t2); i ++)
			{
				bookID = readerBase.Query(PIN)->BorrowBook(rdr->CheckOut());
				bookBase.Query(bookID)->SetState(READER);
				bookBase.Query(bookID)->SetPIN(PIN);
				cout << i + 1 << "\t" << bookBase.Query(bookID)->GetName() << endl;
			}
		}

		Date().Display();

		pushAction(pReaderMenu);
	}
	void readerQueryBook() {
		char name[40];

		cout << "Input a book name (part):";
		cin.ignore();
		cin.get(name, 40, '\n');
		if (bookBase.QueryName(name) != NULL)
		{
			bookBase.QueryName(name)->ShowData();
		}

		pushAction(pReaderMenu);
	}
	void readerListBook() {
		readerBase.Query(PIN)->ShowData();

		pushAction(pReaderMenu);
	}
	void readerCart() {
		rdr->ShowCart();

		pushAction(pReaderMenu);
	}
	void readerExit() {
		pushAction(pMainMenu);
	}
};

// 主函数
int main() {
	LibApp libApp;
	libApp.run();
}