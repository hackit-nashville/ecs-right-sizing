import json
import pandas as pd
from enum import Enum


class Interval(Enum):
    HOUR = "HOUR",
    DOW = "DOW",
    DOW_HOURLY = "DOW_HOURLY"


def calculate(data, interval):

    temp = json.loads(data)

    new_dict = temp['MetricDataResults']
    df = pd.DataFrame.from_dict(new_dict)
    timestamps = df.Timestamps[0]
    task_count = df.Values[0]
    memory_reserved = df.Values[1]
    memory_utilized = df.Values[2]

    data = pd.DataFrame(
        {'Timestamps': timestamps,
         'RunningTaskCount': task_count,
         'MemoryReserved': memory_reserved,
         'MemoryUtilized': memory_utilized
         }
    )

    data['Timestamps'] = pd.to_datetime(data['Timestamps'])

    new_df = data.drop(columns=['RunningTaskCount', 'MemoryReserved'])

    if interval.toUpperCase() == Interval.HOUR:

        new_df['Hour'] = new_df['Timestamps'].apply(lambda x: x.hour)
        final_df = new_df.groupby(['Hour'])['MemoryUtilized'].mean().reset_index()
        return final_df

    elif interval.toUpperCase() == Interval.DOW:

        new_df['DOW'] = new_df['Timestamps'].apply(lambda x: x.dayofweek)
        final_df = new_df.groupby(['DOW'])['MemoryUtilized'].mean().reset_index()
        return final_df

    elif interval.toUpperCase() == Interval.DOW_HOURLY:

        new_df['Hour'] = new_df['Timestamps'].apply(lambda x: x.hour)
        new_df['DOW'] = new_df['Timestamps'].apply(lambda x: x.dayofweek)
        final_df = new_df.groupby(['DOW',' Hour'])['MemoryUtilized'].mean().reset_index()
        return final_df

    else:
        print("Not a valid interval. Please use: HOUR, DOW, or DOW_HOURLY")
        return None
