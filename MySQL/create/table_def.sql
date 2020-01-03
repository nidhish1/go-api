

CREATE TABLE Lynk_Journey (
	JourneyId varchar(30)  NOT NULL,
						   OrderRequestID varchar(18) ,
						   DriverId varchar(30) ,
						   StartDriverGeo varchar(60) ,
						   EndDriverGeo varchar(60) ,
						   Status varchar(10) ,
						   AcceptTime datetime ,
						   ArrivalTime datetime ,
						   StartTime datetime ,
						   ReachTime datetime ,
						   EndTime datetime ,
						   PRIMARY KEY (JourneyId),
						   FOREIGN KEY (OrderRequestID) REFERENCES Lynk_OrderRequest(TranId)
) 





CREATE TABLE Lynk_OrderRequest (
	TranId varchar(18) not null ,
								CustomerId varchar(30)  ,
								SourceLatLng varchar(60)    ,
								DestLatLng varchar(60)  ,
								TruckType varchar(10) ,
								GoodsType varchar(10) ,
								RequestedDate datetime ,
								Status varchar(10) ,
								
								CONSTRAINT PRIMARY KEY (TranId)
) 



CREATE TABLE Lynk_JourneyTrack (
	TrackId varchar(50)  NOT NULL,
								JourneyID varchar(30)  NULL,
								Lat varchar(30) NULL,
								Lng varchar(30)   NULL,
								ReceivedTime datetime ,
								UpdatedTime datetime ,
								alt varchar(30)  ,
								accuracy varchar(30)  NULL,
								altaccuracy varchar(30)  NULL,
								heading varchar(30)  NULL,
								speed varchar(30)  NULL,
								PRIMARY KEY (TrackId),
								FOREIGN KEY (JourneyID) REFERENCES Lynk_Journey(JourneyId)
) 





